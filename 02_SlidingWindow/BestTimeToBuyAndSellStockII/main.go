package main

func maxProfit(prices []int) int {
	if len(prices) == 1 {
		return 0
	}

	min := prices[0]
	profit := 0
	length := len(prices)

	for r := 0; r < length; r++ {
		if r == length-1 {
			profit += prices[r] - min
			break
		}
		if prices[r] > prices[r+1] {
			if prices[r] > min {
				profit += prices[r] - min
			}
			min = prices[r+1]
		}
	}

	return profit
}

// []int{7, 1, 5, 3, 6, 4}, 7
// двигать Р в случае, если Р больше чем Р - 1

// []int{7, 1, 5, 3, 4, 4}, 7
// min = 3, prices[r] = 6
