package main

func characterReplacement(s string, k int) int {
	freqMap := make(map[byte]int)

	res, maxFreq, left, right := 0, 0, 0, 0

	for ; right < len(s); right++ {
		freqMap[s[right]]++
		if v := freqMap[s[right]]; v > maxFreq {
			maxFreq = v
		}

		for right-left+1-maxFreq > k {
			freqMap[s[left]]--
			left++
		}

		if right-left+1 > res {
			res = right - left + 1
		}
	}

	return res
}

// func characterReplacement(s string, k int) int {
// 	maxLength := 0
// 	var (
// 		left         int
// 		right        int
// 		extraByte    byte
// 		hasExtraByte bool
// 	)

// 	windowContent := make(map[byte]int)

// 	for ; right < len(s); right++ {
// 		if _, ok := windowContent[s[right]]; ok {
// 			if hasExtraByte {
// 				if extraByte == s[right] {
// 					windowContent[s[right]]++
// 				} else {
// 					if len(windowContent) == k+1 {

// 						// reduceWindow
// 						for hasExtraByte {
// 							b := s[left]
// 							if windowContent[b] == 1 {
// 								delete(windowContent, b)
// 							} else {
// 								windowContent[b]--
// 								if windowContent[b] < 2 {
// 									hasExtraByte = false
// 								}
// 							}
// 							left++
// 						}
// 					}
// 				}
// 			} else {
// 				hasExtraByte = true
// 				extraByte = s[right]
// 				windowContent[s[right]]++
// 			}
// 		} else {
// 			if len(windowContent) == k+1 {
// 				// reduce window
// 				if windowContent[s[left]] == 1 {
// 					delete(windowContent, s[left])
// 				} else {
// 					windowContent[s[left]]--
// 				}
// 				left++
// 			}
// 			windowContent[s[right]]++
// 		}
// 		if maxLength < right-left+1 {
// 			maxLength = right - left + 1
// 		}
// 	}

// 	return maxLength
// }

// DHNSDKLDLSDNJKNN
