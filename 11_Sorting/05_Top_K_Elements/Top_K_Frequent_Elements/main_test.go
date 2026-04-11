package main

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		input    []int
		k        int
		expected []int
	}{
		{[]int{1, 1, 1, 2, 2, 3}, 2, []int{1, 2}},
		{[]int{1}, 1, []int{1}},
		{[]int{1, 5, 5, 5, 5, 5, 3, 3, 3, 3}, 1, []int{5}},
		{[]int{1, 2, 1, 2, 3, 3}, 1, []int{1}},
		{[]int{1, 2, 1, 2, 1, 2, 3, 1, 3, 2}, 2, []int{1, 2}},
		{[]int{0, -1, -2, -5, 1, -10, -1000000, 100, 100, 100, 100, 99, 99, 99, 0}, 3, []int{100, 99, 0}},
	}

	for _, tt := range tests {
		output := topKFrequent(tt.input, tt.k)

		if !reflect.DeepEqual(output, tt.expected) {
			t.Errorf("got %v, need %v", output, tt.expected)
		}
	}
}
