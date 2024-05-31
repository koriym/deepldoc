package main

import (
	"deepl/translator"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"unicode"
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

	lines := strings.Split(string(fileContent), "\n")
	paragraphs := processParagraphs(lines)
	paragraphText := strings.Join(paragraphs, "\n")

	wrapped := wrapCodeBlocks(paragraphText)

	translatedContent, err := translator.Translate(wrapped, targetLang)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		return
	}

	restoredContent := removeIgnoreTags(translatedContent)

	relativePath := strings.TrimPrefix(path, directory+string(os.PathSeparator))
	newPath := filepath.Join(filepath.Dir(directory), targetLang, relativePath)

	newDir := filepath.Dir(newPath)
	os.MkdirAll(newDir, 0755)

	err = ioutil.WriteFile(newPath, []byte(restoredContent), 0644)
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

func isAlnum(s string) bool {
	return regexp.MustCompile(`\w`).MatchString(s)
}

func isBlockDelimiter(line string) bool {
	trimmedLine := strings.TrimSpace(line)
	isDelimeter := trimmedLine == "```" || trimmedLine == "~~~" || trimmedLine == "<ignode>" || trimmedLine == "</ignode>"

	return isDelimeter
}

func processParagraphs(lines []string) []string {

	var paragraphs []string
	var tempLine []string
	var inBlock bool // Records whether the line you are currently looking at is in a block

	for _, line := range lines {
		if isBlockDelimiter(line) {
			if len(tempLine) > 0 {
				paragraphs = append(paragraphs, strings.Join(tempLine, " "))
				tempLine = nil
			}
			paragraphs = append(paragraphs, line)
			inBlock = !inBlock
		} else if inBlock {
			paragraphs = append(paragraphs, line)
		} else if len(strings.TrimFunc(line, unicode.IsSpace)) == 0 || !isAlnum(line) {
			if len(tempLine) > 0 {
				paragraphs = append(paragraphs, strings.Join(tempLine, " "))
				tempLine = nil
			}
			if isAlnum(line) {
				tempLine = append(tempLine, line)
			} else {
				paragraphs = append(paragraphs, line)
			}
		} else {
			tempLine = append(tempLine, line)
		}
	}
	if len(tempLine) > 0 {
		paragraphs = append(paragraphs, strings.Join(tempLine, " "))
	}

	return paragraphs
}

func wrapCodeBlocks(input string) string {
	codeBlockRegex := regexp.MustCompile("(```[\\s\\S]*?```|`.*?`|~~~[\\s\\S]*?~~~|\\[[^\\]]*\\]\\([^\\)]*\\))")
	matches := codeBlockRegex.FindAllString(input, -1)
	for _, match := range matches {
		placeholder := "<ignore>" + match + "</ignore>"
		input = strings.Replace(input, match, placeholder, 1)
	}
	return input
}

func removeIgnoreTags(input string) string {
	// Regular expressions that matches <ignore> and </ignore>
	startTag := regexp.MustCompile("<ignore>")
	endTag := regexp.MustCompile("</ignore>")

	// Replace tags with an empty string
	input = startTag.ReplaceAllString(input, "")
	input = endTag.ReplaceAllString(input, "")

	return input
}
