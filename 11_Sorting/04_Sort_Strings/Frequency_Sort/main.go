package main

import (
	"sort"
)

func frequencySort(s string) string {
	charMap := make(map[byte]int)
	res := []byte(s)

	for _, b := range res {
		charMap[b]++
	}

	sort.Slice(res, func(i, j int) bool {
		if charMap[res[i]] == charMap[res[j]] {
			return res[i] < res[j]
		}
		return charMap[res[i]] > charMap[res[j]]
	})

	return string(res)
}
