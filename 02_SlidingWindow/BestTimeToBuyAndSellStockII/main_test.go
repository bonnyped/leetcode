package main

import "testing"

func TestMaxProfitII(t *testing.T) {
	tests := []struct {
		prices   []int
		expected int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 7},
		{[]int{1, 2, 3, 4, 5}, 4},
		{[]int{7, 6, 4, 3, 1}, 0},
		{[]int{15, 17, 0, 1, 0, 5}, 8},
		{[]int{1, 10, 9, 20}, 20},
		{[]int{2}, 0},
		{[]int{30, 32, 100, 129}, 99},
	}
	for _, tt := range tests {
		output := maxProfit(tt.prices)

		if output != tt.expected {
			t.Errorf("test case: %v, got %d, need %d", tt.prices, output, tt.expected)
		}
	}
}

// ограничения 1 <= длина массива <= 3* 10^4
// 0 <= price <= 10^4
// дивгаться во времени можно только вперед
