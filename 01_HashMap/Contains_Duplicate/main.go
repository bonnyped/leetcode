package main

func containsNearbyDuplicate(nums []int, k int) bool {
	search := make(map[int]int)

	for idx, num := range nums {
		if i, found := search[num]; found && abs(i-idx) <= k {
			return true
		}
		search[num] = idx
	}
	return false
}

func abs(i int) int {
	if i < 0 {
		i *= -1
	}
	return i
}

// func containsDuplicate(nums []int) bool {
// 	accumulator := make(map[int]struct{})

// 	for _, v := range nums {
// 		if _, hasNum := accumulator[v]; hasNum {
// 			return true
// 		}
// 		accumulator[v] = struct{}{}
// 	}

// 	return false
// }
