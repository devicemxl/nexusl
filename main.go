// main.go
package main

import (
	"fmt"
)

func main() {
	input := `$david do:"corre" how:"rapido";`
	//input := `$david has:(this do:"corre" how:"rapido")`
	/*
		$robot move:forward when:"now";
		$object inspect:detail_level;
		manzana es "roja";` // Keeping a flat triplet for testing both
	*/
	l := NewLexer(input)
	p := NewParser(l)

	program := p.ParseProgram()

	if len(p.Errors()) != 0 {
		fmt.Println("Parser Errors:")
		for _, msg := range p.Errors() {
			fmt.Printf("\t%s\n", msg)
		}
		return
	}

	fmt.Println("Successfully parsed program:")
	fmt.Println(program.String())
}
