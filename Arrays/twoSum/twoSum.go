func twoSum(nums []int, target int) []int {
	indexMap := make(map[int]int)
	for i, num := range nums {
		if _, found := indexMap[target-num]; found {
			return []int{indexMap[target-num], i}
		} else {
			indexMap[num] = i
		}
	}
	return []int{}
}