func longestConsecutive(nums []int) int {
	maxSeq, currSeq := 0, 0
	set := arrayToSet(nums)
	for _, num := range nums {
		currSeq = 0
		if set[num-1] {
			continue
		}
		currSeq += 1
		currNum := num + 1
		for set[currNum] {
			currSeq += 1
			currNum += 1
		}
		if currSeq > maxSeq {
			maxSeq = currSeq
		}
	}
	if currSeq > maxSeq {
		maxSeq = currSeq
	}
	return maxSeq
}

func arrayToSet(nums []int) map[int]bool {
	set := map[int]bool{}
	for _, num := range nums {
		if !set[num] {
			set[num] = true
		}
	}
	return set
}