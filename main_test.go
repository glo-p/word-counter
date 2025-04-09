package main

import (
	"testing"
)

func TestCountWords(t *testing.T) {

	content := []byte("Hello world")
	expected := 2
	actual := countWords(content)
	if actual != expected {
		t.Fail()
		t.Errorf("Expected %d, but got %d", expected, actual)
		//panic("Test failed")
	}
}

func TestCountWordsEmpty(t *testing.T) {
	content := []byte("")
	expected := 0
	actual := countWords(content)
	if actual != expected {
		t.Fail()
		t.Errorf("Expected %d, but got %d", expected, actual)
	}
}

func TestCountWordsMultipleSpaces(t *testing.T) {
	content := []byte("Hello   world")
	expected := 2
	actual := countWords(content)
	if actual != expected {
		t.Errorf("Expected %d, but got %d", expected, actual)
	}
}

func TestCountWordsLeadingSpaces(t *testing.T) {
	content := []byte("   Hello world")
	expected := 2
	actual := countWords(content)
	if actual != expected {
		t.Errorf("Expected %d, but got %d", expected, actual)
	}
}
