package main

import "fmt"

func main() {
	fmt.Println(containsDuplicate([]int{1, 2, 3, 1}))
}

func containsDuplicate(nums []int) bool {
	accumulator := make(map[int]struct{})

	for v := range nums {
		if _, hasNum := accumulator[v]; hasNum {
			return true
		}
		accumulator[v] = struct{}{}
	}

	return false
}
