package main

import (
	"reflect"
	"testing"
)

func TestWrapCodeBlocks(t *testing.T) {
	var tests = []struct {
		input    string
		expected string
	}{
		{
			"Here is my code: ```go\nfmt.Println(\"Hello, world!\")\n``` Cool, right?",
			"Here is my code: <ignore>```go\nfmt.Println(\"Hello, world!\")\n```</ignore> Cool, right?",
		},
		{
			"No code here!",
			"No code here!",
		},
		{
			"```python\ndef hello():\n    print(\"Hello, world!\")\n```\n```ruby\nputs \"Hello, world!\"\n```",
			"<ignore>```python\ndef hello():\n    print(\"Hello, world!\")\n```</ignore>\n<ignore>```ruby\nputs \"Hello, world!\"\n```</ignore>",
		},
		{
			"# ActiveRecord\n\nHello I am\nWorld!\n\n```php+json\n// no translate!\n```\n\n```php\nabc\ndef\nghi\n```\n\n```php\n// no translate!\n```",
			"# ActiveRecord\n\nHello I am\nWorld!\n\n<ignore>```php+json\n// no translate!\n```</ignore>\n\n<ignore>```php\nabc\ndef\nghi\n```</ignore>\n\n<ignore>```php\n// no translate!\n```</ignore>",
		},
	}

	for _, tt := range tests {
		testname := tt.input
		t.Run(testname, func(t *testing.T) {
			ans := wrapCodeBlocks(tt.input)
			if ans != tt.expected {
				t.Errorf("got %s, want %s", ans, tt.expected)
			}
		})
	}
}

func TestProcessParagraphs(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []string
	}{
		{
			name: "Test Case 1",
			input: []string{
				"This is the first line of the first paragraph.",
				"This is the second line of the first paragraph.",
				"",
				"This is the first line of the second paragraph.",
			},
			expected: []string{
				"This is the first line of the first paragraph. This is the second line of the first paragraph.",
				"",
				"This is the first line of the second paragraph.",
			},
		},
		{
			name: "Test Case 2",
			input: []string{
				"This is a single line paragraph.",
			},
			expected: []string{
				"This is a single line paragraph.",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := processParagraphs(tc.input)
			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("got %v, want %v", got, tc.expected)
			}
		})
	}
}
