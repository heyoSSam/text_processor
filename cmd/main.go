package main

import (
	"log"
	"os"
)

func main() {
	inputFile, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	text := string(inputFile)
	// proccessed(text)

	err = os.WriteFile(os.Args[2], []byte(text), 0644)
	if err != nil {
		log.Fatal(err)
	}

	//add timer
}
