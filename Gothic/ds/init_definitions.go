// ds/init_definitions.go
package ds

import (
	"database/sql"
	"fmt"
	"log"
	"strconv" // Necesario para convertir string a float
	"strings" // Necesario para strings.Fields

	_ "github.com/mattn/go-sqlite3" // Importar el driver de SQLite
)

// parseEmbeddingString convierte un string de números separados por espacios a []float32.
func parseEmbeddingString(s string) ([]float32, error) {
	if s == "" {
		return nil, nil
	}
	// Divide el string por espacios y filtra los campos vacíos
	parts := strings.Fields(s) // strings.Fields maneja múltiples espacios y recorta

	embedding := make([]float32, len(parts))
	for i, part := range parts {
		f, err := strconv.ParseFloat(part, 32) // Parsear a float32
		if err != nil {
			return nil, fmt.Errorf("failed to parse float '%s': %w", part, err)
		}
		embedding[i] = float32(f)
	}
	return embedding, nil
}

// LoadSystemDefinitionsFromDB carga los símbolos del sistema desde una DB SQLite.
// dbPath es la ruta al archivo definitions.db
func LoadSystemDefinitionsFromDB(dbPath string) error {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return fmt.Errorf("failed to open system definitions DB at %s: %w", dbPath, err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT public_name, thing, embedding_data FROM system_symbols")
	if err != nil {
		return fmt.Errorf("failed to query system symbols from DB: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var name, thingStringFromDB string // Cambiamos el nombre para evitar confusión
		var embeddingDataString sql.NullString
		if err := rows.Scan(&name, &thingStringFromDB, &embeddingDataString); err != nil {
			return fmt.Errorf("failed to scan symbol row: %w", err)
		}

		s := NewSymbol()
		s.AssignPublicName(name)

		// ¡AQUÍ ESTÁ EL CAMBIO CLAVE!
		// Convertir el string leído de la DB a ds.ThingType
		s.SetThing(ThingType(thingStringFromDB)) // Casteo explícito a ThingTyp

		// Parsear el string de embedding
		if embeddingDataString.Valid && embeddingDataString.String != "" {
			embedding, err := parseEmbeddingString(embeddingDataString.String)
			if err != nil {
				log.Printf("Warning: Failed to parse embedding string for symbol %s: %v", name, err)
			} else {
				s.Embedding = embedding
			}
		}

		// Aquí es donde asignarías los Proc a las macros/funciones built-in
		switch name {
		case "fact":
			s.Proc = func(args ...interface{}) (interface{}, error) {
				fmt.Printf("Macro 'fact' called with args: %v\n", args)
				return nil, nil
			}
		case "def":
			s.Proc = func(args ...interface{}) (interface{}, error) {
				fmt.Printf("Macro 'def' called with args: %v\n", args)
				return nil, nil
			}
		case "func":
			s.Proc = func(args ...interface{}) (interface{}, error) {
				fmt.Printf("Macro 'func' called with args: %v\n", args)
				return nil, nil
			}

		case "macroA": // CASO PARA TEST
			s.Proc = func(args ...interface{}) (interface{}, error) {
				fmt.Printf("Macro 'macroA' called with args: %v\n", args)
				return nil, nil
			}
		}
	}

	fmt.Println("System definitions loaded successfully into memory, including embeddings (parsed from string).")
	return nil
}
