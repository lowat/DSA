func isAnagram(s string, t string) bool {
	sCounts := charCounts(s)
	tCounts := charCounts(t)
	for i, _ := range sCounts {
		if sCounts[i] != tCounts[i] {
			return false
		}
	}
	return true
}

func charCounts(word string) [26]int {
	result := [26]int{}
	for _, c := range word {
		index := int(rune(c) - rune('a'))
		result[index] = result[index] + 1
	}
	return result
}