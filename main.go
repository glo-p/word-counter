package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func countWords(content []byte) int {
 
	if( len(content) == 0) {
		return 0
	}

	words := strings.Split(string(content), ` `)
	for i := 0; i < len(words); i++ {
		if words[i] == "" {
			words = slices.Delete(words, i, i+1)
			i--
		}
	}
	return len(words)
}

func main() {
	fmt.Println("Words Counter")
	content, error := os.ReadFile("words.txt")

	if error != nil {
		fmt.Println("Error:", error)
		return
	}

	fmt.Println("Total words:", countWords(content))
}