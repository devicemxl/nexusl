package main

import lx "github.com/device/nexsusL/pkg/lexer"

func main() {
	line := "Hello, world! This is a test.\n Let's see how it works."
	words := lx.StepOne(line)
	for _, word := range words {
		println(word)
	}
}
