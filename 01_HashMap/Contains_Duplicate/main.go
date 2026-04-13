package main

func containsDuplicate(nums []int) bool {
	accumulator := make(map[int]struct{})

	for _, v := range nums {
		if _, hasNum := accumulator[v]; hasNum {
			return true
		}
		accumulator[v] = struct{}{}
	}

	return false
}
