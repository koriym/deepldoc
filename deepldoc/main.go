package main

import (
	"deepl/translator"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
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
	trimmed := strings.TrimSpace(line)
	// check if line is a block delimiter
	return strings.HasPrefix(trimmed, "`") || strings.HasPrefix(trimmed, "~")
}

func processParagraphs(lines []string) []string {
	var paragraphs []string
	var tempLine []string
	var inBlock bool // track whether the current line is in a block

	for _, line := range lines {
		if isBlockDelimiter(line) {
			inBlock = !inBlock
			if len(tempLine) > 0 {
				paragraphs = append(paragraphs, strings.Join(tempLine, " "))
				tempLine = nil
			}
			paragraphs = append(paragraphs, line)
		} else {
			if !inBlock {
				if len(strings.TrimSpace(line)) == 0 {
					if len(tempLine) > 0 {
						paragraphs = append(paragraphs, strings.Join(tempLine, " "))
						tempLine = nil
					}
					paragraphs = append(paragraphs, line)
				} else {
					tempLine = append(tempLine, line)
				}
			} else {
				paragraphs = append(paragraphs, line)
			}
		}
	}
	if len(tempLine) > 0 {
		paragraphs = append(paragraphs, strings.Join(tempLine, " "))
	}

	return paragraphs
}

func wrapCodeBlocks(input string) string {
	codeBlockRegex := regexp.MustCompile("(```[\\s\\S]*?```|`.*?`|~~~[\\s\\S]*?~~~|\\[[^\\]]*\\]\\([^\\)]*\\))")
	matches := codeBlockRegex.FindAllStringIndex(input, -1)
	offset := 0
	for _, match := range matches {
		start, end := match[0]+offset, match[1]+offset
		original := input[start:end]
		placeholder := "<ignore>" + original + "</ignore>"
		input = input[:start] + placeholder + input[end:]
		offset += len(placeholder) - len(original)
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
