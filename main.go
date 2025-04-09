package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Words Counter")
	content, error := os.ReadFile("words.txt")

	if error != nil {
		fmt.Println("Error:", error)
		return
	}
	words := string(content)
	wordsCount := len(words)
	fmt.Println("Total words:", wordsCount)
}