package main

import "testing"

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		prices   []int
		expected int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{7, 6, 4, 3, 1}, 0},
		{[]int{0, 2, 3, 10, 16}, 16},
		{[]int{5, 10, 3, 9}, 6},
		{[]int{5, 7, 10, 8, 20}, 15},
	}
	for _, tt := range tests {
		output := maxProfit(tt.prices)

		if output != tt.expected {
			t.Errorf("Test case %v: got %d, need %d", tt.prices, output, tt.expected)
		}
	}
}

// constraints 1 <= len(prices) < 10^5 (100 000)
//  0 <= price <= 10^4 (10 000)
