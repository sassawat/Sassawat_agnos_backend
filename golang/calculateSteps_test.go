package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestCalculateSteps(t *testing.T) {
	tests := []struct {
		name     string
		password string
		expected int
	}{
		{"TooShort", "aA1", 3},
		{"StrongEnough", "1445D1cd", 0},
		{"AllLower", "abcde", 2},
		{"AllUpper", "ABCDE", 2},
		{"RepeatingCharacters", "aaaBBB111", 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			steps := calculateSteps(tt.password)
			assert.Equal(t, tt.expected, steps)
		})
	}
}
