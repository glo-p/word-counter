package main

import (
	"testing"
)

type TestCase struct {
	content  []byte
	expected int
}

func TestCountWords(t *testing.T) {
	testCases := []TestCase{
		{[]byte("Hello world"), 2},
		{[]byte("Hello   world"), 2},
		{[]byte("   Hello world"), 2},
		{[]byte("Hello world "), 2},
		{[]byte("Hello"), 1},
		{[]byte(""), 0},
	}

	for _, testCase := range testCases {
		actual := countWords(testCase.content)
		if actual != testCase.expected {
			t.Errorf("Expected %d, but got %d", testCase.expected, actual)
		}
	}
}
