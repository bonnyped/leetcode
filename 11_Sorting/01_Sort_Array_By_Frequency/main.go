package main

import (
	"sort"
)

// func frequencySort(nums []int) []int {
// 	mp := make(map[int]int)
// 	for _, n := range nums {
// 		mp[n]++
// 	}
// 	sort.SliceStable(nums, func(a, b int) bool {
// 		if mp[nums[a]] != mp[nums[b]] {
// 			return mp[nums[a]] < mp[nums[b]]
// 		}
// 		return nums[a] > nums[b]
// 	})
// 	return nums
// }

func frequencySort(nums []int) []int {
	result := []int{}
	numFreq := make(map[int]int)
	sameFreq := make(map[int][]int)

	for _, v := range nums {
		numFreq[v]++
	}

	keys := []int{}
	for key, value := range numFreq {
		sameFreq[value] = append(sameFreq[value], key)
	}

	for key := range sameFreq {
		keys = append(keys, key)
		if len(sameFreq[key]) > 1 {
			sort.Slice(sameFreq[key], func(i, j int) bool {
				return sameFreq[key][i] > sameFreq[key][j]
			})
		}
	}

	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	for _, keyForSameFreq := range keys {
		keysForNumFreq := sameFreq[keyForSameFreq]
		for i := 0; i < len(keysForNumFreq); i++ {
			currValue := keysForNumFreq[i]
			for j := 0; j < numFreq[currValue]; j++ {
				result = append(result, currValue)
			}
		}

	}

	return result
}
