package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Words Counter")
	content, error := os.ReadFile("words.txt")

	if error != nil {
		fmt.Println("Error:", error)
		return
	}
	words := strings.Split(string(content), ` `)
	wordsCount := len(words)
	
	fmt.Println("Total words:", wordsCount)
}