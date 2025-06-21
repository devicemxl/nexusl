package main

import (
	"bytes"
	"context"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io" // Para leer la respuesta HTTP
	"log"
	"math"
	"net/http" // ¡Este es el paquete clave para la llamada HTTP directa!
	"os"
	"sort"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"go.etcd.io/bbolt"
)

// Chunk representa una unidad de texto procesada y fragmentada de la base de conocimientos de NexusL.
// Cada Chunk es un segmento de información autónomo diseñado para ser fácilmente buscable
// y recuperable en el contexto de un sistema RAG (Retrieval-Augmented Generation).
// Estos chunks son la 'memoria' textual de la que el sistema RAG extrae información
// para responder preguntas o guiar el razonamiento de los agentes LLM.
type Chunk struct {
	// ID es el identificador único para este chunk. Se utiliza como clave
	// tanto para almacenar el contenido del chunk como su embedding vectorial
	// en bboltDB, permitiendo una rápida recuperación.
	ID string `json:"id"`

	// Content almacena el texto real del fragmento de información.
	// Es el contenido que se utiliza para generar el embedding y es el texto
	// que se recupera para ser pasado como contexto a un LLM.
	Content string `json:"content"`

	// SectionHeading proporciona contexto adicional sobre el origen del chunk,
	// como el título de la sección o documento del que fue extraído.
	// Esto puede ser útil para la trazabilidad o para mejorar la comprensión
	// por parte del LLM si se incluye en el prompt.
	SectionHeading string `json:"section_heading"`
}

// SearchResult representa un único resultado obtenido de una operación de búsqueda de similitud.
// Contiene la identificación del chunk que fue encontrado como relevante y una puntuación
// que indica qué tan similar es a la consulta original. Estos resultados son el paso
// intermedio entre encontrar los embeddings más cercanos y recuperar el contenido textual
// completo para ser utilizado por un LLM o un agente de NexusL.
type SearchResult struct {
	// ID es el identificador único del Chunk que fue encontrado como relevante.
	// Este ID se usa posteriormente para recuperar el contenido completo del Chunk
	// desde bboltDB.
	ID string

	// Score representa la puntuación de similitud (usualmente similitud de coseno)
	// entre el embedding de la consulta y el embedding del Chunk.
	// Un valor más alto indica una mayor similitud semántica.
	// Los resultados se ordenan por esta puntuación en orden descendente.
	Score float64
}

// EmbeddingFunc defines a function signature for generating vector embeddings from text.
// This abstraction allows the system to be flexible and swap between different
// embedding models and providers (e.g., OpenAI, VoyageAI, Google Gemini)
// without changing the core logic of how chunks are ingested or searched.
//
// Usage and Logic:
// This function type serves as a contract for any embedding provider.
// It takes a Go context.Context for managing deadlines and cancellations,
// and the input 'text' (typically a Chunk.Content or a user query) that needs to be embedded.
// It returns a slice of float32 representing the high-dimensional vector embedding,
// or an error if the embedding generation fails.
//
// Functions/Processes Involved:
//   - IngestData: This function utilizes an EmbeddingFunc to convert the
//     'Content' of each Chunk into its corresponding embedding vector, which is
//     then stored in the bboltDB.
//   - main (for query embedding): When a user submits a search query, an
//     EmbeddingFunc is used to generate the embedding for that query, allowing
//     it to be compared against the stored embeddings in the knowledge base.
//   - LLM Integration (future): This function will also be crucial for embedding
//     user queries before performing RAG (Retrieval-Augmented Generation) to
//     retrieve relevant context for an LLM.
type EmbeddingFunc func(ctx context.Context, text string) ([]float32, error)

// OpenAIEmbeddingRequest defines the structure for the JSON payload sent to the OpenAI Embeddings API.
// This struct is manually crafted based on OpenAI's official API documentation,
// allowing for direct HTTP communication without relying on a third-party SDK.
//
// Logic and Purpose:
// It encapsulates the necessary parameters required by the OpenAI Embeddings endpoint
// to generate a vector representation of input text. By defining this structure,
// the application can easily marshal (convert to JSON) Go data into the format
// expected by the OpenAI API.
//
// Fields:
//   - Input: A slice of strings representing the text(s) for which embeddings are to be generated.
//     For this RAG system, it typically contains a single string: either a chunk's content
//     during ingestion or a user's query during a search operation.
//   - Model: The identifier of the OpenAI embedding model to use (e.g., "text-embedding-3-small").
//     This specifies which specific AI model will process the input text.
//   - EncodingFormat: An optional field to specify the desired encoding format for the output embeddings.
//     "float" is used here to receive the embeddings as an array of floating-point numbers,
//     which is suitable for cosine similarity calculations.
//
// Functions/Processes Involved:
//   - NewEmbeddingFuncOpenAI: This function uses OpenAIEmbeddingRequest to construct
//     the body of the HTTP POST request sent to the OpenAI API endpoint for embeddings.
//     It marshals an instance of this struct into a JSON byte slice before sending.
type OpenAIEmbeddingRequest struct {
	Input          []string `json:"input"`
	Model          string   `json:"model"`
	EncodingFormat string   `json:"encoding_format,omitempty"` // "float" or "base64"
}

// OpenAIEmbeddingResponse defines the structure for the JSON response received from the OpenAI Embeddings API.
// This struct is manually defined based on OpenAI's official API documentation, enabling direct
// HTTP communication without relying on a dedicated SDK.
//
// Logic and Purpose:
// It's designed to unmarshal (convert from JSON) the API's response into a Go-friendly format.
// This allows the application to easily access the generated embeddings and other metadata
// returned by the OpenAI service.
//
// Fields:
//   - Data: A slice of anonymous structs, where each element represents an embedding for one of the
//     input texts. In this RAG system's current use case, this slice will typically contain
//     a single element, as we send one text (a chunk or a query) per request.
//   - Embedding: A slice of float64 representing the high-dimensional vector embedding itself.
//     OpenAI returns embeddings as float64, which are then converted to float32 for internal
//     consistency and memory efficiency within the RAG system.
//   - Index: The index of the input text within the request batch. For single-text requests,
//     this will typically be 0.
//   - Object: The type of object returned, usually "embedding".
//   - Model: The identifier of the OpenAI model that was used to generate the embeddings (e.g., "text-embedding-3-small").
//     This confirms which model processed the request.
//   - Usage: Contains information about the token consumption for the request.
//   - PromptTokens: The number of tokens in the input text(s).
//   - TotalTokens: The total number of tokens processed for the request (usually same as PromptTokens for embeddings).
//   - Object: The type of the top-level API response object, typically "list".
//
// Functions/Processes Involved:
//   - NewEmbeddingFuncOpenAI: This function receives the raw HTTP response from the OpenAI API,
//     reads its body, and then unmarshals the JSON content into an instance of OpenAIEmbeddingResponse.
//     It then extracts the `Embedding` field from the `Data` slice, converts it to `[]float32`,
//     and returns it for use in similarity calculations.
type OpenAIEmbeddingResponse struct {
	Data []struct {
		Embedding []float64 `json:"embedding"` // OpenAI returns float64
		Index     int       `json:"index"`
		Object    string    `json:"object"`
	} `json:"data"`
	Model string `json:"model"`
	Usage struct {
		PromptTokens int `json:"prompt_tokens"`
		TotalTokens  int `json:"total_tokens"`
	} `json:"usage"`
	Object string `json:"object"`
}

// NewEmbeddingFuncOpenAI creates an EmbeddingFunc that directly interacts with the OpenAI Embeddings API
// via HTTP requests. This function serves as the concrete implementation for generating embeddings
// when OpenAI is the chosen provider, adhering to the EmbeddingFunc interface. By using direct HTTP
// calls, it offers maximum flexibility and avoids dependency on a specific OpenAI Go SDK,
// making it easier to switch to other embedding providers (e.g., VoyageAI, Google Gemini) in the future.
//
// Parameters:
//   - apiKey: The API key required for authenticating requests to the OpenAI API.
//   - model: The identifier of the specific OpenAI embedding model to use (e.g., "text-embedding-3-small").
//   - APIURL: The full URL endpoint for the OpenAI Embeddings API (e.g., "https://api.openai.com/v1/embeddings").
//     Passing this as a parameter enhances flexibility, allowing easy switching between different
//     OpenAI endpoints or even custom proxy endpoints if needed.
//
// Logic and Purpose:
// This function returns an anonymous function (an implementation of EmbeddingFunc) that encapsulates
// the entire HTTP request-response cycle for generating an embedding.
//  1. It configures an `http.Client` with a timeout for robust network operations.
//  2. Inside the returned function, it constructs an `OpenAIEmbeddingRequest` with the provided text, model,
//     and desired encoding format ("float").
//  3. The request payload is marshaled into JSON.
//  4. An `http.Request` is created, including the `Context` for cancellation and the necessary
//     `Content-Type` and `Authorization` headers.
//  5. The request is sent, and the HTTP response is read.
//  6. It explicitly checks for non-200 HTTP status codes, indicating API errors, and returns a detailed error.
//  7. The JSON response is unmarshaled into an `OpenAIEmbeddingResponse` struct.
//  8. It validates that a valid embedding was received.
//  9. Finally, it converts the `[]float64` embedding received from OpenAI into `[]float32` for
//     consistency with the rest of the RAG system's internal representation.
//
// Functions/Processes Involved:
//   - Initialization in main: This function is called once during the application's setup
//     to create the `embeddingFunc` instance that will be used throughout the program's lifecycle.
//   - IngestData: During the ingestion process, `IngestData` repeatedly calls the `EmbeddingFunc`
//     (which is the function returned by `NewEmbeddingFuncOpenAI`) to generate embeddings
//     for each chunk of text before storing them in bboltDB.
//   - Query Processing: When a user query is received, the same `EmbeddingFunc` is invoked
//     to generate the embedding for the query, enabling semantic search against the stored embeddings.
//   - LLM Integration: It forms the critical first step in the RAG pipeline for LLM interactions,
//     converting natural language queries into a vector space for retrieval.
func NewEmbeddingFuncOpenAI(apiKey, model string, APIURL string) EmbeddingFunc {
	// You can use an http.Client for more control (timeouts, keep-alives)
	client := &http.Client{
		Timeout: 30 * time.Second, // Timeout for the HTTP call
	}
	openaiAPIURL := APIURL

	return func(ctx context.Context, text string) ([]float32, error) {
		reqBody := OpenAIEmbeddingRequest{
			Input:          []string{text},
			Model:          model,
			EncodingFormat: "float", // Request float format
		}

		jsonBody, err := json.Marshal(reqBody)
		if err != nil {
			return nil, fmt.Errorf("error marshalling JSON request: %w", err)
		}

		// Create the HTTP POST request
		httpReq, err := http.NewRequestWithContext(ctx, "POST", openaiAPIURL, bytes.NewBuffer(jsonBody))
		if err != nil {
			return nil, fmt.Errorf("error creating HTTP request: %w", err)
		}

		// Add required headers
		httpReq.Header.Set("Content-Type", "application/json")
		httpReq.Header.Set("Authorization", "Bearer "+apiKey)

		// Perform the HTTP call
		httpResp, err := client.Do(httpReq)
		if err != nil {
			return nil, fmt.Errorf("error performing HTTP call to OpenAI: %w", err)
		}
		defer httpResp.Body.Close()

		// Read the response
		respBytes, err := io.ReadAll(httpResp.Body)
		if err != nil {
			return nil, fmt.Errorf("error reading HTTP response from OpenAI: %w", err)
		}

		// Handle API errors (status codes != 2xx)
		if httpResp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("OpenAI API error (code %d): %s", httpResp.StatusCode, string(respBytes))
		}

		// Unmarshal the JSON response
		var openAIResp OpenAIEmbeddingResponse
		if err := json.Unmarshal(respBytes, &openAIResp); err != nil {
			return nil, fmt.Errorf("error unmarshalling OpenAI JSON response: %w", err)
		}

		// Verify that a valid embedding was received
		if len(openAIResp.Data) == 0 || len(openAIResp.Data[0].Embedding) == 0 {
			return nil, fmt.Errorf("no valid embedding received from OpenAI for text: '%s'", text)
		}

		// Convert []float64 to []float32 for internal use
		embedding64 := openAIResp.Data[0].Embedding
		embedding32 := make([]float32, len(embedding64))
		for i, v := range embedding64 {
			embedding32[i] = float32(v)
		}
		return embedding32, nil
	}
}

// cosineSimilarity calculates the cosine similarity between two given vectors (v1 and v2).
// Cosine similarity is a metric used to measure how similar two non-zero vectors are,
// indicating the cosine of the angle between them. A value of 1 means the vectors are
// perfectly similar (pointing in the same direction), 0 means they are orthogonal
// (no similarity), and -1 means they are perfectly dissimilar (pointing in opposite directions).
// In the context of embeddings, it quantifies the semantic similarity between two pieces of text.
//
// Parameters:
// - v1: The first vector, represented as a slice of float32.
// - v2: The second vector, represented as a slice of float32.
//
// Logic and Purpose:
// The function computes the cosine similarity using the formula:
// Cosine Similarity(A, B) = (A . B) / (||A|| * ||B||)
// Where:
// - (A . B) is the dot product of vectors A and B.
// - ||A|| is the Euclidean norm (magnitude) of vector A.
// - ||B|| is the Euclidean norm (magnitude) of vector B.
//
// 1. **Input Validation:** It first checks for invalid input conditions:
//   - If either vector is empty, or if their lengths do not match, it returns 0.0,
//     as similarity cannot be meaningfully calculated.
//
// 2. **Dot Product and Norms Calculation:** It iterates through the vectors to compute:
//   - `dotProduct`: The sum of the products of corresponding elements of v1 and v2.
//   - `normA`: The sum of the squares of elements in v1 (used to calculate ||A||).
//   - `normB`: The sum of the squares of elements in v2 (used to calculate ||B||).
//     3. **Division by Zero Prevention:** Before the final division, it checks if either `normA`
//     or `normB` is zero. If so, it returns 0.0 to prevent division by zero, as a zero-norm
//     vector (a zero vector) has no direction.
//     4. **Final Calculation:** It returns the dot product divided by the product of the
//     Euclidean norms (square roots of `normA` and `normB`).
//
// Functions/Processes Involved:
//   - `findTopKSimilar`: This function is the primary consumer of `cosineSimilarity`.
//     It uses `cosineSimilarity` to compare a query embedding against every embedding
//     in the loaded knowledge base (`allEmbeddings`), thereby quantifying the semantic
//     relevance of each stored chunk to the user's query.
//   - **RAG Search Mechanism:** It is a core component of the Retrieval-Augmented Generation
//     (RAG) search mechanism, enabling the system to identify and rank the most relevant
//     pieces of information from the knowledge base.
//   - **Agent Reasoning (future):** In more advanced NexusL agents, cosine similarity
//     could be used to compare conceptual embeddings to make decisions or draw inferences
//     based on semantic closeness of ideas.
func cosineSimilarity(v1, v2 []float32) float64 {
	if len(v1) == 0 || len(v2) == 0 || len(v1) != len(v2) {
		return 0.0 // Invalid vectors
	}

	dotProduct := 0.0
	normA := 0.0
	normB := 0.0
	for i := 0; i < len(v1); i++ {
		dotProduct += float64(v1[i] * v2[i])
		normA += float64(v1[i] * v1[i])
		normB += float64(v2[i] * v2[i])
	}

	if normA == 0 || normB == 0 {
		return 0.0 // Prevent division by zero if a vector is zero
	}
	return dotProduct / (math.Sqrt(normA) * math.Sqrt(normB))
}

// findTopKSimilar identifies and returns the top K most semantically similar chunks
// from the loaded knowledge base based on a given query embedding. This function
// is central to the retrieval phase of the RAG (Retrieval-Augmented Generation) system,
// allowing it to pinpoint the most relevant information to a user's query.
//
// Parameters:
//   - queryEmbedding: A slice of float32 representing the vector embedding of the user's
//     query or the concept being searched for. This is the "needle" in the haystack.
//   - allEmbeddings: A map where keys are chunk IDs (string) and values are their
//     corresponding vector embeddings ([]float32). This represents the entire
//     knowledge base loaded into memory for quick lookup. This is the "haystack."
//   - k: An integer specifying the maximum number of top similar chunks to return.
//     This controls the size of the context provided to a subsequent LLM or processing step.
//
// Logic and Purpose:
//  1. **Similarity Calculation:** The function iterates through every chunk embedding
//     in the `allEmbeddings` map. For each chunk, it calls `cosineSimilarity` to
//     calculate how semantically similar its embedding is to the `queryEmbedding`.
//     This yields a `score` for each chunk, indicating its relevance.
//  2. **Result Accumulation:** Each calculated score, along with the corresponding
//     chunk `ID`, is stored in a `SearchResult` struct. These `SearchResult` instances
//     are appended to a `results` slice.
//  3. **Sorting by Relevance:** Once all similarities are calculated, the `results` slice
//     is sorted in descending order based on the `Score` field. This places the most
//     relevant chunks at the beginning of the list.
//  4. **Top K Selection:** Finally, the function returns a slice containing only the
//     top `k` elements from the sorted `results`. If the total number of available
//     chunks is less than `k`, all available chunks are returned.
//
// Functions/Processes Involved:
//   - **Core Retrieval Mechanism:** This function is the heart of the RAG system's
//     retrieval component. It directly implements the semantic search functionality.
//   - `cosineSimilarity`: It heavily relies on the `cosineSimilarity` function to
//     quantify the relatedness between embeddings.
//   - **Context for LLMs/Agents:** The IDs of the top K similar chunks are then used
//     to retrieve the actual text content of those chunks from `bboltDB`. This retrieved
//     text forms the "augmented context" that is passed to a Large Language Model (LLM)
//     or a NexusL agent, enabling them to generate more accurate, factual, and
//     contextually relevant responses.
//   - **Chatbot/Agent Response Generation:** In the context of a chatbot or intelligent agent
//     for NexusL, the output of this function directly determines which pieces of your
//     NexusL documentation or knowledge will inform the agent's reply.
func findTopKSimilar(queryEmbedding []float32, allEmbeddings map[string][]float32, k int) []SearchResult {
	var results []SearchResult
	for id, emb := range allEmbeddings {
		score := cosineSimilarity(queryEmbedding, emb)
		results = append(results, SearchResult{ID: id, Score: score})
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].Score > results[j].Score
	})

	if len(results) > k {
		return results[:k]
	}
	return results
}

// Definición de la estructura para la solicitud API desde Streamlit
type QueryRequest struct {
	Query string `json:"query"` // El texto de la pregunta del usuario
}

// Definición de la estructura para la respuesta API hacia Streamlit
type QueryResponse struct {
	Answer string `json:"answer"`          // La respuesta generada por el LLM
	Error  string `json:"error,omitempty"` // Campo opcional para mensajes de error
}

// LLMCompletionRequest define la estructura para la solicitud de "completado de chat" a la API de OpenAI.
// Este struct es crucial para construir el payload JSON que se envía a los endpoints
// de chat de OpenAI, permitiendo una interacción conversacional.
type LLMCompletionRequest struct {
	Model       string    `json:"model"`                 // El modelo de chat a usar (ej., "gpt-3.5-turbo", "gpt-4")
	Messages    []Message `json:"messages"`              // <--- ¡Cambiado aquí!
	Temperature float64   `json:"temperature,omitempty"` // Controla la creatividad (0.0-2.0)
	MaxTokens   int       `json:"max_tokens,omitempty"`  // Límite de tokens en la respuesta
}

// LLMCompletionResponse define la estructura para la respuesta de "completado de chat" de la API de OpenAI.
// Esto permite deserializar la respuesta JSON de OpenAI en un formato Go manejable,
// extrayendo la respuesta generada por el modelo.
type LLMCompletionResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int64  `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Index   int `json:"index"`
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
		LogProbs     interface{} `json:"logprobs"` // Puede ser null
		FinishReason string      `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}

// Message representa un mensaje individual en una conversación con un LLM,
// incluyendo el rol del emisor (ej. "system", "user", "assistant") y el contenido del texto.
// Esta estructura es compatible con las APIs de chat de LLMs como OpenAI y Gemini.
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// callLLM interactúa con la API de OpenAI Chat Completions para obtener una respuesta.
// Esta función es el puente entre tu lógica RAG y el modelo de lenguaje grande externo.
//
// Parámetros:
//   - ctx: Contexto de Go para control de timeout y cancelación.
//   - promptMessages: Slice de structs con {Role, Content} que representan la conversación.
//     Esto incluirá tu prompt enriquecido (contexto RAG + pregunta del usuario).
//   - llmModel: El nombre del modelo de LLM a usar (ej., "gpt-3.5-turbo").
//   - openaiAPIKey: La clave API para autenticación con OpenAI.
//   - openaiChatAPIURL: La URL del endpoint de la API de chat de OpenAI.
//
// Retorna:
//   - La respuesta de texto generada por el LLM.
//   - Un error si la llamada API falla o la respuesta es inválida.
func callLLM(ctx context.Context, promptMessages []Message, llmModel, openaiAPIKey, openaiChatAPIURL string) (string, error) {
	client := &http.Client{Timeout: 60 * time.Second} // Aumentar timeout para LLMs

	reqBody := LLMCompletionRequest{
		Model:       llmModel,
		Messages:    promptMessages,
		Temperature: 0.7, // Valor común para respuestas equilibradas
		MaxTokens:   500, // Límite para la respuesta del LLM
	}

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("error marshalling LLM request JSON: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, "POST", openaiChatAPIURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", fmt.Errorf("error creating LLM HTTP request: %w", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", "Bearer "+openaiAPIKey)

	httpResp, err := client.Do(httpReq)
	if err != nil {
		return "", fmt.Errorf("error performing HTTP call to LLM: %w", err)
	}
	defer httpResp.Body.Close()

	respBytes, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading LLM HTTP response: %w", err)
	}

	if httpResp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("LLM API error (code %d): %s", httpResp.StatusCode, string(respBytes))
	}

	var llmResponse LLMCompletionResponse
	if err := json.Unmarshal(respBytes, &llmResponse); err != nil {
		return "", fmt.Errorf("error unmarshalling LLM JSON response: %w", err)
	}

	if len(llmResponse.Choices) == 0 || llmResponse.Choices[0].Message.Content == "" {
		return "", fmt.Errorf("no valid response content received from LLM")
	}

	return llmResponse.Choices[0].Message.Content, nil
}

// serveQuery maneja las peticiones HTTP entrantes del frontend de Streamlit.
// Esta es la función principal que convierte tu aplicación RAG en un servicio API.
func serveQuery(
	w http.ResponseWriter, // Objeto para escribir la respuesta HTTP
	r *http.Request, // Objeto que representa la petición HTTP entrante
	db *bbolt.DB,
	embeddingFunc EmbeddingFunc,
	allEmbeddings map[string][]float32, // Embeddings cargados en memoria
	openaiAPIKey string,
	llmModel string,
	openaiChatAPIURL string, // ¡Este parámetro ya trae la URL correcta desde main!
) {
	// Asegura que la respuesta sea JSON y permita CORS (si Streamlit está en otro dominio)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*") // Para desarrollo; en producción, restringir.
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Manejar peticiones OPTIONS para CORS pre-flight
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Método no permitido. Solo POST.", http.StatusMethodNotAllowed)
		return
	}

	var req QueryRequest
	// Decodifica el cuerpo JSON de la petición HTTP a la estructura QueryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Error al decodificar la petición JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	userQuery := req.Query
	if userQuery == "" {
		http.Error(w, "La pregunta no puede estar vacía.", http.StatusBadRequest)
		return
	}

	log.Printf("Recibida pregunta de Streamlit: '%s'", userQuery)

	// --- Lógica RAG: Aquí se recupera el contexto relevante ---
	queryEmbedding, err := embeddingFunc(r.Context(), userQuery) // Usar r.Context() del request HTTP
	if err != nil {
		log.Printf("Error al generar embedding para la pregunta: %v", err)
		json.NewEncoder(w).Encode(QueryResponse{Answer: "", Error: "Error interno al procesar la pregunta (embedding)."})
		return
	}

	topK := 2 // Número de chunks más relevantes a recuperar
	results := findTopKSimilar(queryEmbedding, allEmbeddings, topK)

	var contextChunksContent strings.Builder
	bucketChunks := []byte("ChunksBucket")

	err = db.View(func(tx *bbolt.Tx) error {
		bktChunks := tx.Bucket(bucketChunks)
		if bktChunks == nil {
			return fmt.Errorf("ChunksBucket no encontrado en la base de datos")
		}

		for _, res := range results {
			chunkBytes := bktChunks.Get([]byte(res.ID))
			if chunkBytes == nil {
				log.Printf("Advertencia: Chunk con ID %s no encontrado en la base de datos.", res.ID)
				continue
			}
			var chunk Chunk
			if err := json.Unmarshal(chunkBytes, &chunk); err != nil {
				log.Printf("Advertencia: Error al deserializar chunk ID %s: %v", res.ID, err)
				continue
			}
			contextChunksContent.WriteString(chunk.Content)
			contextChunksContent.WriteString("\n\n") // Separador entre chunks
		}
		return nil
	})

	if err != nil {
		log.Printf("Error al recuperar chunks de contexto: %v", err)
		json.NewEncoder(w).Encode(QueryResponse{Answer: "", Error: "Error interno al recuperar contexto."})
		return
	}

	// --- Preparar prompt enriquecido para el LLM ---
	systemPrompt := "Eres un asistente de NexusL. Responde preguntas basándote únicamente en el contexto proporcionado. Si la respuesta no está en el contexto, indica que no tienes la información."
	if contextChunksContent.Len() > 0 {
		systemPrompt += "\n\nContexto:\n" + contextChunksContent.String()
	} else {
		log.Println("No se encontró contexto relevante para la pregunta.")
		// Opcional: Aquí podrías añadir un mensaje de fallback o un contexto predeterminado
	}

	promptMessages := []Message{
		{Role: "system", Content: systemPrompt},
		{Role: "user", Content: userQuery},
	}

	// --- Llamada al LLM ---
	// ¡SE ELIMINA ESTA LÍNEA! CAMBIA EL VALOR--> openaiChatAPIURL = os.Getenv("OPENAI_CHAT_API_URL")

	llmResponse, err := callLLM(
		r.Context(),
		promptMessages,
		llmModel, // <-- Aquí va el nombre del modelo (ej. "gpt-3.5-turbo")
		openaiAPIKey,
		openaiChatAPIURL, // <-- Aquí va la URL completa del API de Chat (ej. "https://api.openai.com/v1/chat/completions")
	)
	if err != nil {
		log.Printf("Error al llamar al LLM: %v", err)
		json.NewEncoder(w).Encode(QueryResponse{Answer: "", Error: "Error al generar respuesta del LLM."})
		return
	}

	log.Printf("Respuesta del LLM: '%s'", llmResponse)

	// --- Enviar respuesta al frontend ---
	json.NewEncoder(w).Encode(QueryResponse{Answer: llmResponse})
}

func main() {
	// --- FASE 1: Configuración Inicial del Sistema RAG ---
	// Carga de Variables de Entorno
	configFile := "/Users/davidochoacorrales/Documents/GitHub/nexusl/rag/env/.env"
	if err := godotenv.Load(configFile); err != nil {
		fmt.Printf("Advertencia: No se pudo cargar el archivo de configuracion: %v\n", err)
	}

	// Validación de API Keys
	openaiAPIKey := os.Getenv("OPENAI_API_KEY")
	if openaiAPIKey == "" {
		log.Fatalf("Error: La variable de entorno OPENAI_API_KEY no está configurada o está vacía.")
	}
	openaiChatAPIURL := os.Getenv("OPENAI_CHAT_API_URL")
	if openaiChatAPIURL == "" {
		log.Fatalf("Error: La variable de entorno OPENAI_CHAT_API_URL no está configurada o está vacía. Ej: https://api.openai.com/v1/chat/completions")
	}

	// Configuración de la Ruta de la Base de Datos bbolt
	dbPath := os.Getenv("NEXUS_KB_PATH")
	if dbPath == "" {
		log.Fatalf("Error: La variable de entorno NEXUS_KB_PATH no está configurada o está vacía.")
	}

	// Apertura/Creación de la Base de Datos bbolt
	bboltDB, err := bbolt.Open(dbPath, 0600, &bbolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatalf("Error al abrir bboltDB: %v", err)
	}
	defer bboltDB.Close()

	// Configuración del Modelo de Embeddings y Creación de la Función de Embedding
	embeddingModel := os.Getenv("EMBEDDING_ENGINE")
	if embeddingModel == "" {
		log.Fatalf("Error: La variable de entorno EMBEDDING_ENGINE no está configurada o está vacía.")
	}
	embeddingFunc := NewEmbeddingFuncOpenAI(openaiAPIKey, embeddingModel, os.Getenv("OPEN_AI_EMBEDDING_URL"))
	// Asegúrate de que OPEN_AI_EMBEDDING_URL también esté en tu .env

	// --- FASE 2: Carga de Embeddings en Memoria para Búsqueda ---
	allEmbeddings := make(map[string][]float32)
	bucketEmbeddings := []byte("EmbeddingsBucket")

	fmt.Println("Cargando embeddings desde bbolt a memoria para búsqueda...")
	err = bboltDB.View(func(tx *bbolt.Tx) error {
		bktEmbeddings := tx.Bucket(bucketEmbeddings)
		if bktEmbeddings == nil {
			fmt.Println("El bucket de embeddings aún no existe. La KB está vacía o no se ha ingerido nada.")
			return nil
		}

		cursor := bktEmbeddings.Cursor()
		for k, v := cursor.First(); k != nil; k, v = cursor.Next() {
			chunkID := string(k)
			var embeddingVector []float32
			buf := bytes.NewReader(v)
			if err := binary.Read(buf, binary.LittleEndian, &embeddingVector); err != nil {
				fmt.Printf("Advertencia: Error al deserializar embedding para ID %s: %v\n", chunkID, err)
				continue
			}
			allEmbeddings[chunkID] = embeddingVector
		}
		return nil
	})

	if err != nil {
		log.Fatalf("Error al cargar embeddings desde bbolt: %v", err)
	}
	fmt.Printf("Carga de embeddings completada. %d embeddings cargados.\n", len(allEmbeddings))

	// --- FASE 3: Configuración y Lanzamiento del Servidor HTTP ---
	// Este es el nuevo bloque que transforma main en un servidor API.
	// Define el endpoint `/query` que Streamlit consumirá.
	http.HandleFunc("/query", func(w http.ResponseWriter, r *http.Request) {
		// Asegúrate de que LLM_MODEL se está pasando como 'llmModel' y OPENAI_CHAT_API_URL como 'openaiChatAPIURL'
		serveQuery(w, r, bboltDB, embeddingFunc, allEmbeddings, openaiAPIKey, os.Getenv("LLM_MODEL"), openaiChatAPIURL) // <--- Esta línea
	})

	port := os.Getenv("GO_API_PORT") // Puerto donde escuchará tu API de Go
	if port == "" {
		port = "8080" // Puerto por defecto si no se especifica
	}
	listenAddr := ":" + port

	fmt.Printf("Servidor Go API escuchando en http://localhost%s\n", listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, nil)) // Inicia el servidor y lo mantiene en ejecución
}
