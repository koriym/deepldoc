package main

import (
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
