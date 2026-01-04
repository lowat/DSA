package main

func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var sCount [26]int
	var tCount [26]int

	for i := 0; i < len(s); i++ {
		sCount[s[i]-'a']++
		tCount[t[i]-'a']++
	}

	for i := 0; i < 26; i++ {
		if sCount[i] != tCount[i] {
			return false
		}
	}

	return true
}
