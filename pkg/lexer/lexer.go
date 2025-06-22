package lexer

import (
	"strings"
)

func StepOne(line string) []string {

	words := strings.Fields(line)

	//fmt.Printf("Original string: %q\n", line)
	//fmt.Printf("Split words: %q\n", words)
	return words
}
