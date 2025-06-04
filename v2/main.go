// main.go
package main

import (
	"fmt"
)

func main() {
	input1 := `$david do:"corre" how:"rapido";`
	// Caso 1: Tripleta calificada con lista de propiedades calificadas
	// $david [do:"corre" do:"baila" do:"vuela"] how:"rapido";
	//input1 := `$david [do:"corre", do:"baila", do:"vuela"] how:"rapido";`
	/*
		// Caso 2: Tripleta calificada con lista de acciones (strings)
		// $david do:["corre", "baila", "vuela"] how:"rapido";
		input2 := `$david do:["corre", "baila", "vuela"] how:"rapido";`

		// Caso 3: Lista de elementos variados
		// $COLECCION is [ "a", "b", "c", 4, $QWE];
		input3 := `$COLECCION is [ "a", "b", "c", 4, $QWE];`

		// Caso 4: Definición de función
		// func:calculaArea param:[ base has Value is int; altura has Value is int; ] code:[ c = a * b; D = c + 1; export: [D]; ];
		input4 := `func:calculaArea: param:[base has Value is int; altura has Value is int;] code:[c = a * b; D = c + 1;] export:[D];`
	*/
	// --- Ejecutar pruebas ---

	fmt.Println("--- Probando Caso 1: Tripleta con lista de propiedades ---")
	l1 := NewLexer(input1) // Usar New para el lexer
	p1 := NewParser(l1)    // Usar NewParser para el parser
	program1 := p1.ParseProgram()
	if len(p1.Errors()) != 0 {
		fmt.Println("Errores del Parser (Caso 1):")
		for _, msg := range p1.Errors() {
			fmt.Printf("\t%s\n", msg)
		}
	} else {
		fmt.Println("Programa parseado exitosamente (Caso 1):")
		fmt.Println(program1.String())
	}
}
