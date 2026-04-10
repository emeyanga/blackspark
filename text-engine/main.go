package main

import (
	"fmt"
	"os"
	"text-engine/internal/tokenizer"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Go usage: go run . input.txt output.txt")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	content, err := os.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}

	result := string(content)

	result = tokenizer.Tokenize(result)

	if err = os.WriteFile(outputFile, []byte(result), 0644); err != nil {
		fmt.Println("Failed to write file!!!", err)
	}
}
