package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	search := [26]int{}

	for i := 0; i < len(s); i++ {
		search[s[i]-'a']++
		search[t[i]-'a']--
	}

	return search == [26]int{}
}
