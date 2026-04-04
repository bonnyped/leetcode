package main

import "fmt"

func main() {
	fmt.Println(characterReplacement("ABABB", 2))
}

// "ABBCALSA"

func characterReplacement(s string, k int) int {
	start, end := 0, 0

	frequency := make(map[byte]int)
	res ,windowLength, maxFrequency := 0, 0

	for ; end < len(s); end++ {
		frequency[s[end]]++ // добавляем в окно букву

		if maxFrequency < frequency[s[end]] {
			maxFrequency = frequency[s[end]]
		}

		if end - start + 1 > k {
			frequency[s[end]]--
			start++
		}

		if

		fmt.Println(v, start)
	}

	return res
	// задача получить самую длинную подстроку с одним символом
	// выявить самую длинную подстроку с одним символорм с К повторением другого символов
	// записать
	// wrong - k >= 0
}


func longestWindow(nums int, condition int) int {
	start, maxLength, res := 0, 0, 0
	
	
	return  res
}