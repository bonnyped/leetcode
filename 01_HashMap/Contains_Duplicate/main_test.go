package main

import (
	"testing"
)

func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		nums     []int
		expected bool
	}{
		{[]int{1, 2, 3, 1}, true},
		{[]int{1, 2, 3, 4}, false},
		{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
		{[]int{}, false},
		{[]int{0}, false},
	}

	for _, tt := range tests {
		result := containsDuplicate(tt.nums)

		if result != tt.expected {
			t.Errorf("got %t, want %t", result, tt.expected)
		}
	}
}
