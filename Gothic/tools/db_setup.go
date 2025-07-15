// db_setup/main.go (Un archivo separado para inicializar la DB)
package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	dbPath := "./db/definitions.db" // Ruta donde se creará la DB

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatalf("Failed to open DB: %v", err)
	}
	defer db.Close()

	// Crear la tabla system_symbols
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS system_symbols (
		public_name TEXT PRIMARY KEY,
		thing TEXT NOT NULL,
		embedding_data TEXT -- Almacena el embedding como un string de floats separados por espacio
	);
	`
	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}
	fmt.Println("Table 'system_symbols' created or already exists.")

	// Insertar datos de ejemplo
	insertDataSQL := `
	INSERT OR REPLACE INTO system_symbols (public_name, thing, embedding_data) VALUES
	('fact', 'TripletScope', '0.1 0.2 0.3'),
	('program', 'TripletScope', '0.4 0.5 0.6'),
	('func', 'TripletScope', '0.7 0.8 0.9'),
	('service', 'TripletScope', '1.0 1.1 1.2'),
	('var', 'TripletScope', '1.3 1.4 1.5'),
	('is', 'Predicate', '0.11 0.22 0.33'),
	('has:', 'Predicate', '0.44 0.55 0.66'),
	('do:', 'Predicate', '0.77 0.88 0.99'),
	('how::', 'Predicate', '1.01 1.12 1.23'),
	('Car', 'Identifier', '0.01 0.02 0.03'),
	('symbol', 'Identifier', '0.04 0.05 0.06');
	`
	_, err = db.Exec(insertDataSQL)
	if err != nil {
		log.Fatalf("Failed to insert data: %v", err)
	}
	fmt.Println("Sample data inserted/replaced.")

	// Opcional: Consulta para verificar los datos
	rows, err := db.Query("SELECT public_name, thing, embedding_data FROM system_symbols")
	if err != nil {
		log.Fatalf("Failed to query data: %v", err)
	}
	defer rows.Close()

	fmt.Println("\n--- Data in system_symbols ---")
	for rows.Next() {
		var name, thing string
		var embedding sql.NullString
		if err := rows.Scan(&name, &thing, &embedding); err != nil {
			log.Fatalf("Failed to scan row: %v", err)
		}
		fmt.Printf("Name: %s, Thing: %s, Embedding: %s\n", name, thing, embedding.String)
	}
}
