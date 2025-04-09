package main

import (
	"testing"
)

type TestCase struct {
	name     string
	content  []byte
	expected int
}

func TestCountWords(t *testing.T) {
	testCases := []TestCase{
		{name: "Simple case", content: []byte("Hello world"), expected: 2},
		{name: "Multiple spaces", content: []byte("Hello   world"), expected: 2},
		{name: "Leading spaces", content: []byte("   Hello world"), expected: 2},
		{name: "Trailing spaces", content: []byte("Hello world "), expected: 2},
		{name: "Single word", content: []byte("Hello"), expected: 1},
		{name: "Empty string", content: []byte(""), expected: 0},
	}

	for _, testCase := range testCases {
		actual := countWords(testCase.content)
		if actual != testCase.expected {
			t.Error(testCase.name)
			t.Errorf("Expected %d, but got %d", testCase.expected, actual)
		}
	}
}
