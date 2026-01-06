package main

func TwoSum(nums []int, target int)[]int {
	prev := make(map[int]int)
	for i, num := range nums {
		if idx, found := prev[target - num]; found {
			return []int{idx, i}
		}
		prev[num] = i
	}
	return []int{}
}
