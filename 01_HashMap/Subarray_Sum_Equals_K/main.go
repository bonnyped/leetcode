package main

func subarraySum(nums []int, k int) int {
	prefixs := map[int]int{0: 1}
	prefixSum := 0
	var result int

	for _, num := range nums {
		prefixSum += num
		if value, hasPrefix := prefixs[prefixSum-k]; hasPrefix {
			result += value
		}
		prefixs[prefixSum]++
	}

	return result
}
