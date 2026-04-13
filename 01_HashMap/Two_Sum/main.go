package main

func twoSum(nums []int, target int) []int {
	res := make([]int, 2)
	search := make(map[int]int)

	for idx, num := range nums {
		if prevIdx, found := search[num]; found {
			res[0] = prevIdx
			res[1] = idx
			break
		}
		search[target-num] = idx
	}

	return res
}
