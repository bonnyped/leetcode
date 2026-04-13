package main

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		input    []int
		target   int
		expected []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
		{[]int{0, 0}, 0, []int{0, 1}},
	}
	for _, tt := range tests {
		output := twoSum(tt.input, tt.target)

		if !reflect.DeepEqual(output, tt.expected) {
			t.Errorf("got %v, need %v", output, tt.expected)
		}
	}
}
