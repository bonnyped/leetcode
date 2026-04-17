package main

import "testing"

func TestCharacterReplacement(t *testing.T) {
	tests := []struct {
		input    string
		k        int
		expected int
	}{
		{"ABAB", 2, 4},
		{"AABABBA", 1, 4},
		{"AABABAABBBBB", 2, 8},
		{"AABBBBBBB", 0, 7},
		{"A", 0, 1},
		{"AAAB", 0, 3},
		{"DHNSDKLDLSDNJKNN", 2, 5},
	}
	for _, tt := range tests {
		output := characterReplacement(tt.input, tt.k)

		if output != tt.expected {
			t.Errorf("Test case %s: got %d, need %d", tt.input, output, tt.expected)
		}
	}
}
