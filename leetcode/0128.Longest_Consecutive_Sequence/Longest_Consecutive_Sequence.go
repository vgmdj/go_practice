package Longest_Consecutive_Sequence

func longestConsecutive(nums []int) int {
	exist, result := make(map[int]bool, 0), 0
	for _, c := range nums {
		exist[c] = true
	}

	for _, c := range nums {
		if !exist[c] {
			continue
		}

		cLen := 1
		for pre := c - 1; exist[pre]; pre-- {
			cLen++
			exist[pre] = false
		}

		for aft := c + 1; exist[aft]; aft++ {
			cLen++
			exist[aft] = false
		}

		if cLen > result {
			result = cLen
		}

	}

	return result

}
