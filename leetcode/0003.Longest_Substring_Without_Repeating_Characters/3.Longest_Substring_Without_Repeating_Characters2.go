package Longest_Substring_Without_Repeating_Characters

func lengthOfLongestSubstring2(s string) int {
	if len(s) == 0 {
		return 0
	}

	position := make([]int, 256)
	for i := range position {
		position[i] = -1
	}

	left, right, max := 0, 0, 1

	for ; right < len(s); right++ {
		if position[s[right]] >= left {
			left = position[s[right]] + 1
		}

		position[s[right]] = right

		if max < right-left+1 {
			max = right - left + 1
		}

		continue

	}

	return max

}
