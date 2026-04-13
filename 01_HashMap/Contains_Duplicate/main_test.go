package main

import (
	"testing"
)

func TestContainsNearbyDuplicate(t *testing.T) {
	tests := []struct {
		input    []int
		k        int
		expected bool
	}{
		{[]int{1, 2, 3, 1}, 3, true},
		{[]int{1, 0, 1, 1}, 1, true},
		{[]int{-1, -10, 0, 1, 1}, 1, true},
		{[]int{1, 2, 3, 1, 2, 3}, 2, false},
		{[]int{1, 2, 3, 1, 2, 3}, 3, true},
		{[]int{}, 0, false},
		{[]int{1, 1}, 1, true},
		{[]int{1, 1}, 0, false},
	}
	for _, tt := range tests {
		output := containsNearbyDuplicate(tt.input, tt.k)
		if output != tt.expected {
			t.Errorf("input %v: got %t, need %t", tt.input, output, tt.expected)
		}

	}

}

// func TestContainsDuplicate(t *testing.T) {
// 	tests := []struct {
// 		nums     []int
// 		expected bool
// 	}{
// 		{[]int{1, 2, 3, 1}, true},
// 		{[]int{1, 2, 3, 4}, false},
// 		{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
// 		{[]int{}, false},
// 		{[]int{0}, false},
// 	}

// 	for _, tt := range tests {
// 		result := containsDuplicate(tt.nums)

// 		if result != tt.expected {
// 			t.Errorf("got %t, want %t", result, tt.expected)
// 		}
// 	}
// }
