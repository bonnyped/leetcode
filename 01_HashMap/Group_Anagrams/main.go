package main

import "maps"

// первое решение
type word struct {
	dic  map[byte]int
	strs []string
}

func groupAnagrams(strs []string) [][]string {
	result := make([]word, 0)

	for _, str := range strs {
		dic := make(map[byte]int)
		for j := 0; j < len(str); j++ {
			dic[str[j]]++
		}

		w := dictionaryExist(result, dic)

		if w == nil {
			result = append(result, word{dic, []string{str}})
		} else {
			w.strs = append(w.strs, str)
		}

	}

	return mergeResults(result)
}

func mergeResults(words []word) [][]string {
	result := make([][]string, 0)

	for _, v := range words {
		result = append(result, v.strs)
	}

	return result
}

func dictionaryExist(result []word, dic map[byte]int) *word {
	for i := range result {
		if maps.Equal(result[i].dic, dic) {
			return &result[i]
		}
	}
	return nil
}
