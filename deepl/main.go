package main

import (
	"deepl/translator"
	"fmt"
	"os"
)

func main() {
	targetLang := "ja"
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./deepl text targetLang")
		os.Exit(1)
	}

	text := os.Args[1]
	if len(os.Args) > 2 {
		targetLang = os.Args[2]
	}

	translatedText, err := translator.Translate(text, targetLang)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%s\n", translatedText)
}
