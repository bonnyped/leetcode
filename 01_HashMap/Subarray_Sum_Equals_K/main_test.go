package main

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		input    []int
		k        int
		expected int
	}{
		{[]int{1, 1, 1}, 2, 2},
		{[]int{1, 2, 3}, 3, 2},
		{[]int{0, 0}, 0, 3},
		{[]int{1, -1, 0}, 0, 3},
		{[]int{3, -95, 92, 0}, 0, 3},
	}
	for _, tt := range tests {
		output := subarraySum(tt.input, tt.k)

		if output != tt.expected {
			t.Errorf("input: %v, got %d, need %d", tt.input, output, tt.expected)
		}
	}
}

// {0:1}
// {0:1, 1:1} p = 0, 0+1 = 1 , r = 0
// {0:1, 1:1, 2:1} p = 1, 1-1 = 0, r = 1
// {0:1, 1:1, 2:1} p = 2, 1+2 = 3, r = 2
