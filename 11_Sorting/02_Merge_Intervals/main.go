package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	result := [][]int{}
	min, max := 0, 0
	length := len(intervals)

	if length <= 1 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	for i := 0; i < length; i++ {
		min = intervals[i][0]
		max = intervals[i][1]
		for j := i + 1; j < length; j++ {
			if intervals[j][0] <= max {
				if intervals[j][1] > max {
					max = intervals[j][1]
				}
				i++
			} else {
				break
			}
		}

		result = append(result, []int{min, max})
	}

	fmt.Println(result)

	return result
}
