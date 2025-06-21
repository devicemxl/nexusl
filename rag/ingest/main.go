// ingest: JSON to embeddings
package main

import (
	"bufio"
	"bytes"
	"context"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io" // Para leer la respuesta HTTP
	"log"
	"net/http" // ¡Este es el paquete clave para la llamada HTTP directa!
	"os"
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

// IngestData processes a file containing raw text chunks, generates their embeddings,
// and stores both the original chunk content and their vector embeddings into a bbolt database.
// This function is crucial for building and updating the knowledge base (KB) that the
// RAG system relies on for information retrieval. It acts as the 'loading dock'
// for new information into the system's long-term memory.
//
// Parameters:
//   - chunksFilePath: The file path to a line-delimited JSON (JSONL) file, where each line
//     represents a `Chunk` object containing text content to be ingested.
//   - bboltDB: A pointer to an open bbolt.DB instance. This is the persistent storage
//     where both the original text chunks and their corresponding embeddings will be saved.
//   - embeddingFunc: An `EmbeddingFunc` that provides the capability to convert text
//     content into numerical vector embeddings. This allows IngestData to be agnostic
//     to the specific embedding model or provider (e.g., OpenAI, VoyageAI, Gemini).
//
// Logic and Purpose:
// The function reads the input `chunksFilePath` line by line. For each valid JSON line
// representing a `Chunk`:
//  1. **Deserialization:** It unmarshals the JSON line into a `Chunk` struct.
//  2. **Storage of Raw Chunk:** The original `Chunk` data (as JSON bytes) is stored
//     in the "ChunksBucket" within `bboltDB`, using the `Chunk.ID` as the key. This ensures
//     that the original textual content can be retrieved later based on its ID.
//  3. **Embedding Generation:** It calls the provided `embeddingFunc` to generate a
//     vector embedding for the `Chunk.Content`. This is a critical step that converts
//     human-readable text into a machine-understandable numerical representation.
//  4. **Serialization and Storage of Embedding:** The generated `[]float32` embedding vector
//     is then serialized into a byte slice using `binary.Write` and stored in the
//     "EmbeddingsBucket" within `bboltDB`, also using the `Chunk.ID` as the key. This
//     allows for efficient retrieval of embeddings during similarity searches.
//
// All database operations are wrapped within a single bbolt `Update` transaction
// to ensure atomicity and data consistency. Error handling includes checks for file
// opening, JSON parsing, embedding generation, and database writes.
//
// Functions/Processes Involved:
//   - **Knowledge Base Creation/Update:** This is the primary function for populating
//     or updating your RAG system's knowledge base. It's designed to be run periodically
//     (e.g., via a cron job or a dedicated update script for NexusL) to keep the
//     information fresh.
//   - **Embedding Generation:** It drives the calls to external embedding APIs
//     (via `embeddingFunc`) for all ingested data.
//   - **Persistent Storage:** It manages the storage of both raw textual data and
//     numerical embeddings in `bboltDB`, making them available for future retrieval operations.
//   - **Separation of Concerns (future):** In a larger application, this function
//     would typically be part of a dedicated 'ingestion service' or 'data loader' component,
//     separate from the query/chatbot interface.
func IngestData(
	chunksFilePath string,
	bboltDB *bbolt.DB,
	embeddingFunc EmbeddingFunc,
) error {
	file, err := os.Open(chunksFilePath)
	if err != nil {
		return fmt.Errorf("error opening chunks file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	chunkCount := 0

	fmt.Println("Starting ingestion process...")

	bucketChunks := []byte("ChunksBucket")
	bucketEmbeddings := []byte("EmbeddingsBucket") // New bucket for embeddings

	err = bboltDB.Update(func(tx *bbolt.Tx) error {
		bktChunks, err := tx.CreateBucketIfNotExists(bucketChunks)
		if err != nil {
			return fmt.Errorf("error creating or opening chunks bucket: %w", err)
		}
		bktEmbeddings, err := tx.CreateBucketIfNotExists(bucketEmbeddings) // Create embeddings bucket
		if err != nil {
			return fmt.Errorf("error creating or opening embeddings bucket: %w", err)
		}

		for scanner.Scan() {
			line := scanner.Text()
			if len(line) == 0 {
				continue
			}

			var chunk Chunk
			if err := json.Unmarshal([]byte(line), &chunk); err != nil {
				fmt.Printf("Warning: invalid line: %v\n", err)
				continue
			}

			chunkBytes, err := json.Marshal(chunk)
			if err != nil {
				fmt.Printf("Error marshalling chunk ID %s: %v\n", chunk.ID, err)
				continue
			}

			// --- Store RAW Chunk in bbolt ---
			if err := bktChunks.Put([]byte(chunk.ID), chunkBytes); err != nil {
				fmt.Printf("Error storing chunk ID %s in bbolt (chunks): %v\n", chunk.ID, err)
				continue
			}

			// --- Generate Embedding ---
			embeddingVector, err := embeddingFunc(context.Background(), chunk.Content)
			if err != nil {
				fmt.Printf("Error generating embedding for chunk ID %s: %v\n", chunk.ID, err)
				continue
			}

			// --- Serialize and Store Embedding in bbolt ---
			var buf bytes.Buffer
			err = binary.Write(&buf, binary.LittleEndian, embeddingVector)
			if err != nil {
				fmt.Printf("Error serializing embedding for chunk ID %s: %v\n", chunk.ID, err)
				continue
			}
			embeddingBytes := buf.Bytes()

			if err := bktEmbeddings.Put([]byte(chunk.ID), embeddingBytes); err != nil {
				fmt.Printf("Error storing embedding for chunk ID %s in bbolt (embeddings): %v\n", chunk.ID, err)
				continue
			}

			chunkCount++
		}
		return nil
	})

	if err != nil {
		return fmt.Errorf("error during bbolt ingestion transaction: %w", err)
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	fmt.Printf("Ingestion completed. Total: %d chunks\n", chunkCount)
	return nil
}

func main() {
	// --- FASE 1: Configuración Inicial del Sistema RAG ---
	// Esta sección se encarga de cargar las variables de entorno y configurar
	// los componentes esenciales del sistema, como la base de datos y el
	// proveedor de embeddings. Es una fase de inicialización crítica
	// que se ejecutaría al inicio de cualquier programa que use el RAG.

	// 1. Carga de Variables de Entorno:
	// Se carga el archivo '.env' para obtener configuraciones sensibles o variables
	// de ruta, como la API Key de OpenAI y las rutas a la base de datos y los chunks.
	// Es crucial para la portabilidad y seguridad de las credenciales.
	configFile := "/Users/davidochoacorrales/Documents/GitHub/nexusl/rag/env/.env"
	if err := godotenv.Load(configFile); err != nil {
		// Se emite una advertencia si el archivo .env no se puede cargar,
		// pero no se detiene la ejecución inmediatamente, ya que algunas
		// variables podrían estar configuradas directamente en el entorno.
		fmt.Printf("Advertencia: No se pudo cargar el archivo de configuracion: %v\n", err)
	}

	// 2. Validación de API Key:
	// Se obtiene la API Key de OpenAI desde las variables de entorno.
	// Si no está configurada, el programa termina, ya que no se pueden
	// generar embeddings sin ella.
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatalf("Error: La variable de entorno OPENAI_API_KEY no está configurada o está vacía.")
	}

	// 3. Configuración de la Ruta de la Base de Datos bbolt:
	// Se obtiene la ruta de la base de datos bbolt donde se almacenará
	// la base de conocimientos (chunks y embeddings).
	dbPath := os.Getenv("NEXUS_KB_PATH")

	// 4. Apertura/Creación de la Base de Datos bbolt:
	// Se abre la base de datos bbolt. Si no existe, se crea.
	// Es la capa de persistencia para el conocimiento del sistema.
	bboltDB, err := bbolt.Open(dbPath, 0600, &bbolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatalf("Error al abrir bboltDB: %v", err)
	}
	// Se asegura que la base de datos se cierre al finalizar la ejecución del programa.
	defer bboltDB.Close()

	// 5. Configuración del Modelo de Embeddings y Creación de la Función de Embedding:
	// Se obtiene el nombre del modelo de embeddings a utilizar (ej. text-embedding-3-small).
	// Se crea la función `embeddingFunc` que encapsula la lógica para llamar
	// a la API de embeddings de OpenAI. Esta función es clave para generar
	// las representaciones vectoriales del texto.
	embeddingModel := os.Getenv("EMBEDDING_ENGINE")
	embeddingWebPath := os.Getenv("OPEN_AI_EMBEDDING_URL")
	embeddingFunc := NewEmbeddingFuncOpenAI(apiKey, embeddingModel, embeddingWebPath) // Pasa la URL de la API

	// --- FASE 2: Carga de Embeddings en Memoria para Búsqueda ---
	// Esta sección se encarga de cargar todos los embeddings existentes desde
	// la base de datos bbolt a la memoria RAM. Esto es necesario para realizar
	// búsquedas de similitud eficientes, ya que comparar vectores en memoria
	// es mucho más rápido que hacerlo directamente desde el disco.
	// Esta fase es parte del inicio de la 'fase de consulta'.

	allEmbeddings := make(map[string][]float32)    // Mapa para almacenar embeddings: ID_Chunk -> Vector_Embedding
	bucketEmbeddings := []byte("EmbeddingsBucket") // Nombre del bucket donde se guardan los embeddings

	fmt.Println("Cargando embeddings desde bbolt a memoria para búsqueda...")
	err = bboltDB.View(func(tx *bbolt.Tx) error {
		bktEmbeddings := tx.Bucket(bucketEmbeddings)
		if bktEmbeddings == nil {
			fmt.Println("El bucket de embeddings aún no existe. La KB está vacía o no se ha ingerido nada.")
			return nil // No es un error si el bucket no existe aún
		}

		cursor := bktEmbeddings.Cursor() // Iterador para recorrer todos los pares clave-valor en el bucket
		for k, v := cursor.First(); k != nil; k, v = cursor.Next() {
			chunkID := string(k)
			var embeddingVector []float32
			// Se lee el vector de bytes y se deserializa de nuevo a []float32
			buf := bytes.NewReader(v)
			err := binary.Read(buf, binary.LittleEndian, &embeddingVector)
			if err != nil {
				fmt.Printf("Advertencia: Error al deserializar embedding para ID %s: %v\n", chunkID, err)
				continue // Continúa con el siguiente embedding si hay un error
			}
			allEmbeddings[chunkID] = embeddingVector // Se guarda en el mapa en memoria
		}
		return nil
	})
	if err != nil {
		log.Fatalf("Error al cargar embeddings desde bbolt a memoria: %v", err)
	}
	fmt.Printf("Cargados %d embeddings en memoria.\n", len(allEmbeddings))

	// --- FASE 3: Ingestión (Actualización) de Datos (Opcional en un proceso de Consulta) ---
	// Esta sección se encarga de leer nuevos chunks desde un archivo JSONL,
	// generar sus embeddings y guardarlos en la base de datos bbolt.

	chunksFile := os.Getenv("CHUNKS_FILE") // Ruta al archivo JSONL con los chunks
	if chunksFile == "" {
		log.Fatalf("Error: La variable de entorno CHUNKS_FILE no está configurada o está vacía.")
	}

	if err := IngestData(chunksFile, bboltDB, embeddingFunc); err != nil {
		log.Fatalf("Error durante la ingestión: %v", err)
	}

}
