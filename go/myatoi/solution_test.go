package myatoi

import (
	"testing"
)

func TestMyAtoi(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "positive number",
			input:    "42",
			expected: 42,
		},
		{
			name:     "negative number",
			input:    "-42",
			expected: -42,
		},
		{
			name:     "with whitespace",
			input:    "   -42",
			expected: -42,
		},
		{
			name:     "with words",
			input:    "4193 with words",
			expected: 4193,
		},
		{
			name:     "words first",
			input:    "words and 987",
			expected: 0,
		},
		{
			name:     "overflow max",
			input:    "2147483648",
			expected: 2147483647,
		},
		{
			name:     "overflow min",
			input:    "-2147483649",
			expected: -2147483648,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.input); got != tt.expected {
				t.Errorf("myAtoi(%q) = %v, want %v", tt.input, got, tt.expected)
			}
		})
	}
}
