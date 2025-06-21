package lexer

import (
	"fmt"
	"strings"
)

func stepOne(line string) []string {

	words := strings.Fields(line)
	fmt.Printf("Original string: %q\n", line)
	fmt.Printf("Split words: %q\n", words)
	return words
}
