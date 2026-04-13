package main

func containsNearbyDuplicate(nums []int, k int) bool {
	search := make(map[int]int)

	for idx, num := range nums {
		if i, found := search[num]; found && idx-i <= k {
			return true
		}
		search[num] = idx
	}
	return false
}
