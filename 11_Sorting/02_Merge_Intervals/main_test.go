package main

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		intervals [][]int
		expected  [][]int
	}{
		{[][]int{{15, 18}, {1, 3}, {2, 6}, {8, 10}}, [][]int{{1, 6}, {8, 10}, {15, 18}}},
		{[][]int{{1, 4}, {4, 5}}, [][]int{{1, 5}}},
		{[][]int{{4, 7}, {1, 4}}, [][]int{{1, 7}}},
		{[][]int{{}}, [][]int{{}}},
		{[][]int{{1, 4}, {2, 3}}, [][]int{{1, 4}}},
	}

	for _, tt := range tests {
		result := merge(tt.intervals)

		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("got %v, want %v", result, tt.expected)
		}
	}
}
