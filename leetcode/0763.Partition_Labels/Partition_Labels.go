package Partition_Labels

func partitionLabels(S string) []int {
	result := make([]int, 0)
	rightBorder := [26]int{}
	for i, c := range S {
		rightBorder[c-'a'] = i
	}

	left, right := 0, 0
	for i, c := range S {
		if rightBorder[c-'a'] > right {
			right = rightBorder[c-'a']
		}

		if i == right {
			result = append(result, right-left+1)
			left = right + 1
		}
	}

	return result
}
