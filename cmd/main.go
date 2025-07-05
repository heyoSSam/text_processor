package main

import (
	"log"
	"os"
	"strings"
	"text_processor/internal/processing"
)

func main() {
	inputFile, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	text := string(inputFile)
	var textTokens []string
	textTokens = processing.Tokenize(text)
	result := strings.Join(processing.Process(textTokens), " ")
	result = processing.FixPunctuation(result)
	err = os.WriteFile(os.Args[2], []byte(result), 0644)
	if err != nil {
		log.Fatal(err)
	}

	//add timer
}
