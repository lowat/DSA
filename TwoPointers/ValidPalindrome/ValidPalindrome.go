import (
	"strings"
)

func isPalindrome(s string) bool {
	cleanedString := cleanString(s)
	left, right := 0, len(cleanedString)-1
	for left < right {
		if cleanedString[left] != cleanedString[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func cleanString(s string) string {
	res := strings.ToLower(s)
	var result strings.Builder
	for i := 0; i < len(res); i++ {
		b := res[i]
		if ('a' <= b && b <= 'z') ||
			('A' <= b && b <= 'Z') ||
			('0' <= b && b <= '9') {
			result.WriteByte(b)
		}
	}
	return result.String()
}