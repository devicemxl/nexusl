// Package main provides the functionality for the pre-ingestion phase of the RAG system.
// This phase is responsible for transforming raw, human-readable text documents (often
// in a custom delimited format) into structured JSON Lines (JSONL) format.
// Each line in the output JSONL file represents a 'chunk' of information,
// ready for the subsequent 'ingest' phase (embedding generation and storage).
package main

import (
	"bufio"         // For efficient reading of files line by line
	"crypto/sha256" // For generating SHA256 hashes to create unique chunk IDs
	"encoding/hex"  // For encoding the SHA256 hash to a hexadecimal string
	"encoding/json" // For marshaling (serializing) Go structs to JSON
	"fmt"           // For formatted I/O (printing messages and errors)
	"log"           // For logging fatal errors
	"os"            // For file system operations (opening, creating files)
	"strings"       // For string manipulation (e.g., checking for substrings, trimming)

	"github.com/joho/godotenv" // For loading environment variables from a .env file
)

// sha256gen generates a SHA256 hash for a given string content.
// This hash is used to create a unique and deterministic ID for each chunk.
// Using a content-based hash ensures that if the chunk content doesn't change,
// its ID remains the same, which is useful for idempotency in ingestion.
//
// Parameters:
//   - content: The string whose SHA256 hash is to be computed. This will typically
//     be the JSON representation of a processed chunk, ensuring the ID reflects
//     all its fields.
//
// Returns:
// - A string representing the hexadecimal encoding of the SHA256 hash.
//
// Usage:
// This function is called by `parseChunks` to assign a unique identifier
// to each structured chunk before it's written to the output JSONL file.
func sha256gen(content string) string {
	hash := sha256.Sum256([]byte(content))
	return hex.EncodeToString(hash[:])
}

// parseChunks reads a source text file, parses it into logical chunks based on
// delimiters and metadata prefixes, and then writes these structured chunks
// into a new JSONL file. This is the core logic of the `preIngest` phase.
//
// The input file is expected to follow a specific custom format with markers
// like "---CHUNK_START---", "source_file:", "content_start", etc.
//
// Parameters:
// - inputPath: The file path to the raw source document that needs to be parsed.
// - outputPath: The file path where the processed JSONL chunks will be written.
//
// Returns:
// - An error if any file operation or parsing fails; otherwise, nil.
//
// Logic and Purpose:
//  1. **File Handling:** Opens the input file for reading and creates/truncates
//     the output file for writing. Ensures files are closed using `defer`.
//  2. **Line-by-Line Scanning:** Uses `bufio.Scanner` to efficiently read the input
//     file line by line.
//  3. **State Machine Parsing:** The `switch` statement acts as a simple state machine
//     to interpret different types of lines:
//     - `---CHUNK_START---`: Signals the beginning of a new chunk, initializing a
//     `processedChunk` map to store its data.
//     - Metadata lines (e.g., `source_file:`, `section_heading:`): Extracts key-value
//     pairs for chunk metadata.
//     - `content_start`: Marks the beginning of the actual text content of the chunk.
//     - Content lines (between `content_start` and `content_end`): Appends these
//     lines to `contentData`.
//     - `content_end`: Signals the end of content, joins `contentData` into a single
//     `content` string, trims whitespace, and assigns it to `processedChunk["content"]`.
//     - `---CHUNK_END---`: Signals the end of a chunk definition. At this point:
//     - The `processedChunk` map is marshaled to JSON.
//     - Its SHA256 hash (based on the entire chunk's JSON representation) is
//     generated using `sha256gen` and assigned as the `id` field.
//     - The final `processedChunk` (including its ID) is stored in `allChunks`,
//     a map using the ID as the key to prevent duplicates if any arise.
//  4. **Output Writing:** After processing all lines in the input file:
//     - It iterates through the `allChunks` map.
//     - Each `chunk` map is marshaled into a compact (non-indented) JSON byte slice.
//     - The JSON data is written to the `outputPath` file, followed by a newline character,
//     ensuring each chunk occupies a single line as required by the JSONL format.
//     - `writer.Flush()` ensures all buffered data is written to the file.
//
// Functions/Processes Involved:
//   - **Data Preparation Pipeline:** This function is the initial step in preparing
//     raw textual knowledge for your RAG system. It transforms unstructured data
//     into the structured `Chunk` format expected by the `ingest` phase.
//   - `sha256gen`: Directly calls `sha256gen` for ID generation.
//   - `json.Marshal`: Utilized for converting Go maps (representing chunks) into JSON strings.
func parseChunks(inputPath, outputPath string) error {
	file, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	allChunks := make(map[string]map[string]interface{}) // Using map[string]interface{} for flexibility in chunk metadata
	var processedChunk map[string]interface{}
	var contentData []string // Temporarily holds lines of content

	for scanner.Scan() {
		line := scanner.Text()

		switch {
		case strings.Contains(line, "ID_AUTO"):
			// Ignore lines containing "ID_AUTO" as they are placeholders in the source format
			continue
		case strings.Contains(line, "---CHUNK_START---"):
			// Mark the beginning of a new chunk, initialize a new map for its data
			processedChunk = make(map[string]interface{})
		case processedChunk != nil && strings.Contains(line, "source_file:"):
			// Extract source file metadata
			processedChunk["source_file"] = strings.TrimSpace(strings.SplitN(line, ":", 2)[1])
		case processedChunk != nil && strings.Contains(line, "section_heading:"):
			// Extract section heading metadata
			processedChunk["section_heading"] = strings.TrimSpace(strings.SplitN(line, ":", 2)[1])
		case processedChunk != nil && strings.Contains(line, "page_number:"):
			// Extract page number metadata
			processedChunk["page_number"] = strings.TrimSpace(strings.SplitN(line, ":", 2)[1])
		case processedChunk != nil && strings.Contains(line, "language:"):
			// Extract language metadata
			processedChunk["language"] = strings.TrimSpace(strings.SplitN(line, ":", 2)[1])
		case processedChunk != nil && strings.Contains(line, "group:"):
			// Extract group metadata
			processedChunk["group"] = strings.TrimSpace(strings.SplitN(line, ":", 2)[1])
		case processedChunk != nil && strings.Contains(line, "tags:"):
			// Extract and parse comma-separated tags
			tags := strings.Split(strings.TrimSpace(strings.SplitN(line, ":", 2)[1]), ",")
			for i := range tags {
				tags[i] = strings.TrimSpace(tags[i])
			}
			processedChunk["tags"] = tags
		case strings.Contains(line, "content_start"):
			// Mark the beginning of content section, initialize content buffer
			contentData = []string{}
		case contentData != nil && !strings.Contains(line, "content_end"):
			// Accumulate content lines
			contentData = append(contentData, line)
		case strings.Contains(line, "content_end"):
			// Mark the end of content, join accumulated lines into single content string
			content := strings.Join(contentData, "")
			processedChunk["content"] = strings.TrimSpace(content)
			contentData = nil // Reset content buffer
		case strings.Contains(line, "---CHUNK_END---"):
			// Mark the end of a chunk. Generate ID and store the complete chunk.
			// Marshal to JSON first to ensure the hash covers all fields for a stable ID.
			chunkJSONBytes, err := json.Marshal(processedChunk)
			if err != nil {
				// Log error but try to continue to process other chunks
				fmt.Printf("Warning: Error marshalling chunk for ID generation: %v\n", err)
				continue
			}
			id := sha256gen(string(chunkJSONBytes))
			processedChunk["id"] = id      // Assign the generated ID
			allChunks[id] = processedChunk // Store chunk in map, using ID as key
			processedChunk = nil           // Reset for next chunk
		}
	}

	// Check for any scanner errors that occurred during scanning
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading input file: %w", err)
	}

	outputFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)
	for _, chunk := range allChunks {
		jsonData, err := json.Marshal(chunk) // Marshal the chunk map into compact JSON
		if err != nil {
			// Log error but try to continue to write other chunks
			fmt.Printf("Warning: Error marshalling chunk for output: %v\n", err)
			continue
		}
		writer.Write(jsonData)   // Write JSON data
		writer.WriteString("\n") // Write newline to ensure JSONL format (one JSON object per line)
	}
	writer.Flush() // Ensure all buffered data is written to the file
	fmt.Printf("Chunks processed and saved to %s\n", outputPath)

	return nil
}

// main function serves as the entry point for the pre-ingestion utility.
// It handles loading environment variables and orchestrates the parsing
// and conversion of the source truth document into JSONL chunks.
//
// This program is designed to be executed manually or as part of an automated
// data preparation pipeline, prior to the embedding ingestion phase.
func main() {
	// --- Phase: Initial Configuration and Environment Setup ---
	// This section handles loading environment variables and validating
	// critical paths needed for the pre-ingestion process.

	// 1. Load Environment Variables:
	// Loads configuration from a '.env' file. This is crucial for managing
	// file paths and other settings without hardcoding them, enhancing
	// portability and security.
	configFile := "/Users/davidochoacorrales/Documents/GitHub/nexusl/rag/env/.env"
	if err := godotenv.Load(configFile); err != nil {
		// Log a warning if the .env file can't be loaded, but proceed as
		// variables might be set directly in the environment.
		fmt.Printf("Warning: Could not load configuration file: %v\n", err)
	}

	// 2. Validate Source and Output Paths:
	// Retrieves the input source file path (`SOURCE_TRUTH`) and the output
	// JSONL chunks file path (`CHUNKS_FILE`) from environment variables.
	// The program will terminate if these critical paths are not configured.
	sourceOfTruthPath := os.Getenv("SOURCE_TRUTH")
	if sourceOfTruthPath == "" {
		log.Fatalf("Error: Environment variable SOURCE_TRUTH is not set or is empty.")
	}
	chunksFilePath := os.Getenv("CHUNKS_FILE")
	if chunksFilePath == "" { // Corrected: should check chunksFilePath here
		log.Fatalf("Error: Environment variable CHUNKS_FILE is not set or is empty.")
	}

	// --- Phase: Chunk Parsing and JSONL Generation ---
	// This is the core operational phase where the raw text is transformed.

	// 3. Execute Chunk Parsing:
	// Calls the `parseChunks` function to perform the main task: reading
	// the source document, parsing it into structured chunks, and writing
	// the output to the specified JSONL file.
	err := parseChunks(
		sourceOfTruthPath,
		chunksFilePath)
	if err != nil {
		fmt.Println("Error during chunk parsing:", err)
		// Optionally, log.Fatalf if this is a critical, non-recoverable error
	} else {
		fmt.Println("Pre-ingestion process completed successfully.")
	}
}
