package main

func main() {
	line := "Hello, world! This is a test.\n Let's see how it works."
	words := lx.stepOne(line)
	for _, word := range words {
		println(word)
	}
}
