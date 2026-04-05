package main

import (
	"reflect"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		strs     []string
		expected [][]string
	}{
		{strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			expected: [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}}},
		{strs: []string{""},
			expected: [][]string{{""}}},
		{strs: []string{"a"},
			expected: [][]string{{"a"}}},
	}

	for _, tt := range tests {
		result := groupAnagrams(tt.strs)
		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("got %v, want%v", result, tt.expected)
		}
	}

}
