package main

import (
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	sortedStrs := make(map[string][]string)

	for _, str := range strs {
		sortedStr := sortString(str)
		sortedStrs[sortedStr] = append(sortedStrs[sortedStr], str)
	}

	return mergedStrs(sortedStrs)
}

func mergedStrs(sortedStrs map[string][]string) [][]string {
	result := make([][]string, len(sortedStrs))
	idx := 0

	for _, value := range sortedStrs {
		result[idx] = append(result[idx], value...)
		idx++
	}

	return result
}

func sortString(str string) string {
	runes := []rune(str)

	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})

	return string(runes)
}

// первое решение
// import "maps"

// type word struct {
// 	dic  map[byte]int
// 	strs []string
// }

// func groupAnagrams(strs []string) [][]string {
// 	result := make([]word, 0)

// 	for _, str := range strs {
// 		dic := make(map[byte]int)
// 		for j := 0; j < len(str); j++ {
// 			dic[str[j]]++
// 		}

// 		w := dictionaryExist(result, dic)

// 		if w == nil {
// 			result = append(result, word{dic, []string{str}})
// 		} else {
// 			w.strs = append(w.strs, str)
// 		}

// 	}

// 	return mergeResults(result)
// }

// func mergeResults(words []word) [][]string {
// 	result := make([][]string, 0)

// 	for _, v := range words {
// 		result = append(result, v.strs)
// 	}

// 	return result
// }

// func dictionaryExist(result []word, dic map[byte]int) *word {
// 	for i := range result {
// 		if maps.Equal(result[i].dic, dic) {
// 			return &result[i]
// 		}
// 	}
// 	return nil
// }
