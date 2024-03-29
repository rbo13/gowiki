package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	expectedStr := "Hello, testing"
	result := hello()

	if result != expectedStr {
		t.Fatalf("Expected %s, got %s", expectedStr, result)
	}
}
