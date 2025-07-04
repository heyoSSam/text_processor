package main

import (
	"fmt"
	"log"
	"os"
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
	result := processing.Process(textTokens)

	fmt.Println(result)
	//err = os.WriteFile(os.Args[2], []byte(text), 0644)
	//if err != nil {
	//	log.Fatal(err)
	//}

	//add timer
}
