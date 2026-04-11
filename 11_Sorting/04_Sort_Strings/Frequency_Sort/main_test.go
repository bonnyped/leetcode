package main

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{"tree", []string{"eert", "eetr"}},
		{"cccaaa", []string{"cccaaa", "aaaccc"}},
		{"Aabb", []string{"bbAa", "bbaA"}},
		{"loveleetcode", []string{"eeeeoollvtdc", "eeeeoollvtcd", "eeeeoolltvdc", "eeeelloocdtv"}},
	}

	for _, tt := range tests {
		output := frequencySort(tt.input)
		if !testAllExpected(output, tt.expected) {
			t.Errorf("got %s, want one of %v", output, tt.expected)
		}
	}
}

func testAllExpected(output string, s []string) bool {
	for _, str := range s {
		if output == str {
			return true
		}
	}
	return false
}
