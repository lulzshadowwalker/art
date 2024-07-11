package util_test

import (
	"testing"

	"github.com/lulzshadowwalker/art/internal/util"
)

func TestUpperFirst(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "single character",
			input:    "a",
			expected: "A",
		},
		{
			name:     "multiple characters",
			input:    "hello",
			expected: "Hello",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := util.UpperFirst(tt.input)
			if actual != tt.expected {
				t.Errorf("expected %q but got %q", tt.expected, actual)
			}
		})
	}
}

func TestLowerFirst(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "single character",
			input:    "A",
			expected: "a",
		},
		{
			name:     "multiple characters",
			input:    "Hello",
			expected: "hello",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := util.LowerFirst(tt.input)
			if actual != tt.expected {
				t.Errorf("expected %q but got %q", tt.expected, actual)
			}
		})
	}
}
