// Gothic/main.go (Actualizado)
package main

import (
	"fmt"

	"github.com/devicemxl/nexusl/ds"                  // Para llamar a LoadSystemDefinitionsFromDB
	"github.com/devicemxl/nexusl/internal/Gothic/ast" // Asegúrate de importar ast
	"github.com/devicemxl/nexusl/internal/Gothic/lexer"
	"github.com/devicemxl/nexusl/internal/Gothic/metamodel"
	"github.com/devicemxl/nexusl/internal/Gothic/parser"
	// Asegúrate de importar token
)

func main() {
	dbPath := "./db/definitions.db" // Asegúrate de que esta ruta sea correcta

	// 1. Cargar las definiciones del sistema desde la DB
	err := ds.LoadSystemDefinitionsFromDB(dbPath)
	if err != nil {
		fmt.Printf("Error loading system definitions from DB: %v\n", err)
		return
	}
	fmt.Println("System definitions loaded from DB.")

	// 2. Crear el facade del metamodelo que usará los símbolos cargados
	mm := metamodel.NewMetamodelFacade()

	input := `fact Car is symbol;` // Tu entrada de prueba

	l := lexer.New(input)
	p := parser.New(l, mm) // Pasa el facade del metamodelo al parser

	program := p.ParseProgram()

	if len(p.Errors()) != 0 {
		fmt.Println("Parser errors:")
		for _, msg := range p.Errors() {
			fmt.Printf("  %s\n", msg)
		}
		return
	}

	fmt.Println("Parsing successful!")
	fmt.Printf("Parsed Program:\n%s\n", program.String())

	// Verificación manual del AST (para depuración)
	if len(program.Statements) > 0 {
		if factStmt, ok := program.Statements[0].(*ast.FactStatement); ok {
			fmt.Printf("  Fact Statement: %s %s %s;\n",
				factStmt.Subject.TokenLiteral(),
				factStmt.Predicate.TokenLiteral(),
				factStmt.Object.TokenLiteral())
			fmt.Printf("  Scope: %s (Type: %s)\n", factStmt.Scope.PublicName, factStmt.Scope.Thing)
		}
	}
}
