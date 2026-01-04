package main

func ContainsDuplicate(nums []int)bool {
	seen := make(map[int]bool)

	for _, num := range nums {
		if _, found := seen[num]; found {
			return true
		}
		seen[num] = true
	}
	return false
}
