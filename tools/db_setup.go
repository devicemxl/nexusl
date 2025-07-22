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
	-- TripletScope: Elementos de la estructura del programa / paradigmas
	('fact',    'TripletScope', '0.10 0.20 0.30'), -- Hecho: una verdad básica
	('rule',    'TripletScope', '0.12 0.22 0.32'), -- Regla: implicación lógica (condición -> consecuencia)
	('logic',   'TripletScope', '0.15 0.25 0.35'), -- Scope de lógica (engloba fact, rule)
	('def',     'TripletScope', '0.40 0.50 0.60'), -- Definición general
	('func',    'TripletScope', '0.42 0.52 0.62'), -- Función (definición de procedimiento o metodo)
	('expr',    'TripletScope', '0.45 0.55 0.65'), -- Expresión simbólica (para cálculos, transformaciones)
	('macro',   'TripletScope', '0.48 0.58 0.68'), -- Macro (expansión de código)
	('program', 'TripletScope', '0.80 0.85 0.90'), -- Scope de programa (nivel superior)

	-- TripletScope: Declaración de Variables
	('var',     'TripletScope', '1.30 1.40 1.50'), -- Declaración de variable mutable
	('let',     'TripletScope', '1.32 1.42 1.52'), -- Declaración de variable de con type inmutable
	('const',   'TripletScope', '1.35 1.45 1.55'), -- Declaración de constante (inmutable global)

	-- Predicate: Verbos y preguntas
	('is',      'Predicate',    '0.11 0.22 0.33'), -- Predicado de igualdad/relación
	('has',     'Predicate',    '0.44 0.55 0.66'), -- Predicado de posesión/propiedad
	('do',      'Predicate',    '0.77 0.88 0.99'), -- Predicado de acción
	('how',     'Predicate',    '1.01 1.12 1.23'), -- manera
	('where',   'Predicate',    '1.03 1.14 1.25'), -- lugar
	('when',    'Predicate',    '1.05 1.16 1.27'), -- tiempo

	-- Identifier: Básico
	('symbol',  'Identifier',   '0.04 0.05 0.06'); -- Un identificador genérico
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
