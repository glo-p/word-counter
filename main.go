package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func fasterCountWords(content []byte) int {
	words := strings.Fields(string(content))
	return len(words)
}

func main() {
	fmt.Println("Words Counter")
	content, error := os.ReadFile("wordsh.txt")

	if error != nil {
		log.Fatalln("Error:", error)	
	}

	fmt.Println("Total words:", fasterCountWords(content))
}