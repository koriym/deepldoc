package main

import (
	"testing"
)

func TestReplaceCodeBlocks(t *testing.T) {
	testCases := []struct {
		name         string
		input        string
		expectedText string
	}{
		{
			name:         "single code block",
			input:        "```go\nfmt.Println(\"Hello, world!\")\n```",
			expectedText: "<ignore lang=\"go\">\nfmt.Println(\"Hello, world!\")\n</ignore>",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actualText := replaceCodeBlocks(tc.input)
			if actualText != tc.expectedText {
				t.Errorf("expected %s, got %s", tc.expectedText, actualText)
			}
		})
	}
}

func TestRestoreCodeBlocks(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "single code block",
			input:    "<ignore lang=\"go\">\nfmt.Println(\"Hello, world!\")\n</ignore>",
			expected: "```go\nfmt.Println(\"Hello, world!\")```",
		},
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := restoreCodeBlocks(tc.input)
			if actual != tc.expected {
				t.Errorf("expected %s, got %s", tc.expected, actual)
			}
		})
	}
}
