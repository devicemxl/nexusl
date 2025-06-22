package main

import (
	"bufio"
	"fmt"
	"strings"

	lx "github.com/devicemxl/nexusl/pkg/lexer"
)

func main() {
	multilineString := ` David is symbol;
	// Hello, world! This is a comment.\n Let's "see how" it works.
	Hola,
Este es un string
que ocupa varias líneas.// este es un comentario too
Incluso puedes incluir
  espacios al principio de las líneas.`
	//
	// Create a new scanner to read the string line by line
	// strings.NewReader converts the string into an io.Reader,
	// which bufio.Scanner can then read from.
	lineNumber := 1
	scanner := bufio.NewScanner(strings.NewReader(multilineString))
	// Step Ones
	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()
		//var theText = line
		var dx = lx.StepOne(line)
		fmt.Printf("Split words: %q\n", dx)
	}

}
