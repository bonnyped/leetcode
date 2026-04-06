package main

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		testCase []int
		expexted []int
	}{
		{[]int{2, 3, 1, 3, 2}, []int{1, 3, 3, 2, 2}},
		{[]int{1, 1, 2, 2, 2, 3}, []int{3, 1, 1, 2, 2, 2}},
		{[]int{-1, 1, -6, 4, 5, -6, 1, 4, 1}, []int{5, -1, 4, 4, -6, -6, 1, 1, 1}},
		{[]int{}, []int{}},
	}

	for _, tt := range tests {
		result := frequencySort(tt.testCase)
		if !reflect.DeepEqual(result, tt.expexted) {
			t.Errorf("got %v, got %v", result, tt.expexted)
		}
	}
}
