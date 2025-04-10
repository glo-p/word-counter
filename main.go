package main

import (
	"fmt"
	"os"
	"strings"
)

func fasterCountWords(content []byte) int {
	words := strings.Fields(string(content))
	return len(words)
}

func main() {
	fmt.Println("Words Counter")
	content, error := os.ReadFile("words.txt")

	if error != nil {
		fmt.Fprintln(os.Stderr, "Error reading file:", error)
		fmt.Println("Error:", error)
		os.Exit(1)	
	}

	fmt.Println("Total words:", fasterCountWords(content))
}