package main

import (
	"sort"
)

func topKFrequent(nums []int, k int) []int {
	compact := []int{}
	freqMap := make(map[int]int)

	for _, num := range nums {
		if _, has := freqMap[num]; !has {
			compact = append(compact, num)
		}
		freqMap[num]++
	}

	sort.Slice(compact, func(i, j int) bool {
		return freqMap[compact[i]] > freqMap[compact[j]]
	})

	return compact[0:k]
}
