package main

import "testing"

func TestValidAnagram(t *testing.T) {
	tests := []struct {
		strS     string
		strT     string
		expected bool
	}{
		{"anagram", "nagaram", true},
		{"rat", "car", false},
		{"ar", "car", false},
	}
	for _, tt := range tests {
		output := isAnagram(tt.strS, tt.strT)

		if output != tt.expected {
			t.Errorf("test case: %s vs %s, got %t, need %t", tt.strS, tt.strT, output, tt.expected)
		}
	}
}
