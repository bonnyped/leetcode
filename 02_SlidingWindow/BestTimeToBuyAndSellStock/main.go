package main

func maxProfit(prices []int) int {
	l, r, profit := 0, 1, 0
	min := prices[l]
	length := len(prices)

	for ; r < length; r++ {
		if prices[r] < min {
			min = prices[r]
			l = r
		}
		if profit < prices[r]-prices[l] {
			profit = prices[r] - prices[l]
		}
	}

	return profit
}

// 100, 200, 10, 190
