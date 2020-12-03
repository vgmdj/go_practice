package Maximum_Repeating_Substring

func maxRepeating(sequence string, word string) int {
	result := 0
	dp := make([]int, len(sequence))
	for i := 0; i < len(sequence)-len(word)+1; i++ {
		if sequence[i:i+len(word)] == word {
			if i >= len(word) {
				dp[i] = dp[i-len(word)] + 1

			} else {
				dp[i] = 1

			}

			result = max(result, dp[i])
		}

	}

	return result

}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b

}
