package main

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"", 0},
	}

	for _, tt := range tests {
		output := lengthOfLongestSubstring(tt.input)

		if output != tt.expected {
			t.Errorf("test case: %s, got %d, need %d", tt.input, output, tt.expected)
		}
	}
}
