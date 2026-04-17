package main

func lengthOfLongestSubstring(s string) int {
	ascii := [128]int{}
	res, left, right := 0, 0, 0

	for ; right < len(s); right++ {
		ascii[s[right]]++

		for ascii[s[right]] > 1 {
			ascii[s[left]]--
			left++
		}

		if right-left+1 > res {
			res = right - left + 1
		}
	}

	return res
}
