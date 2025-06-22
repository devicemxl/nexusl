package main

import (
	"bufio"
	"fmt"
	"strings"

	lx "github.com/devicemxl/nexusl/pkg/lexer"
	tk "github.com/devicemxl/nexusl/pkg/token"
)

// lexer as map
// number of line and tokens
var ThisScript map[int][]tk.Token // Declares a map with string keys and int values

func main() {
	// ¡Aquí está la solución! Inicializa el mapa antes de usarlo.
	ThisScript = make(map[int][]tk.Token)

	multilineString := ` David is symbol;
	// Hello, world! This is a comment.\n Let's "see how" it works.
	Hola,
	/*
Este es un "comentario"
que ocupa varias líneas.// este es un comentario too
*/
"Incluso puedes incluir"
  espacios al principio de las líneas.
  Aqui, /* ESTE TAMBIEN ES COMETARIO */`
	//
	// Create a new scanner to read the string line by line
	// strings.NewReader converts the string into an io.Reader,
	// which bufio.Scanner can then read from.
	lineNumber := 0
	scanner := bufio.NewScanner(strings.NewReader(multilineString))
	// Step Ones
	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()
		var dx = lx.StepOne(line)
		ThisScript[lineNumber] = lx.StepTwo(dx)
		//fmt.Printf("Split words: %q\n", dx)
		fmt.Printf("%v %v\n", lineNumber, ThisScript[lineNumber])
	}
}
