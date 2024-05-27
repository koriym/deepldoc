package main

import (
	"deepl/translator"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	directory := "."
	if len(os.Args) > 1 {
		directory = os.Args[1]
	}

	targetLang := "ja"
	if len(os.Args) > 2 {
		targetLang = os.Args[2]
	}

	extension := ".md"
	if len(os.Args) > 3 {
		extension = "." + os.Args[3] // Add dot to extension
	}

	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			if filepath.Ext(path) == extension { // File matches desired extension
				translateAndSaveFile(path, directory, targetLang)
			} else { // File does not match desired extension, so copy it
				copyFile(path, directory, targetLang)
			}
		}

		return nil
	})

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func translateAndSaveFile(path, directory, targetLang string) {
	fileContent, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		return
	}

	translatedContent, err := translator.Translate(string(fileContent), targetLang)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		return
	}

	relativePath := strings.TrimPrefix(path, directory+string(os.PathSeparator))
	newPath := filepath.Join(filepath.Dir(directory), targetLang, relativePath)

	newDir := filepath.Dir(newPath)
	os.MkdirAll(newDir, 0755)

	err = ioutil.WriteFile(newPath, []byte(translatedContent), 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		return
	}
}

func copyFile(path, directory, targetLang string) {
	output, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		return
	}

	relativePath := strings.TrimPrefix(path, directory+string(os.PathSeparator))
	newPath := filepath.Join(filepath.Dir(directory), targetLang, relativePath)

	newDir := filepath.Dir(newPath)
	os.MkdirAll(newDir, 0755)

	err = ioutil.WriteFile(newPath, output, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		return
	}
}
