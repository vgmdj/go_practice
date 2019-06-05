package Word_Break

func wordBreak(s string, wordDict []string) bool {
	visited := make([]bool, len(s))

	return helper(s, wordDict, 0, visited)
}

func helper(s string, wordDict []string, index int, visited []bool) bool {
	if index >= len(s) {
		return true
	}

	if visited[index] {
		return false
	}

	for _, word := range wordDict {
		if index+len(word) > len(s) || s[index:index+len(word)] != word {
			continue
		}

		if helper(s, wordDict, index+len(word), visited) {
			return true
		}
	}

	visited[index] = true

	return false

}

func wordBreak2(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 0; i < len(dp); i++ {
		if !dp[i] {
			continue
		}

		for _, word := range wordDict {
			if i+len(word) <= len(s) && s[i:i+len(word)] == word {
				dp[i+len(word)] = true
			}
		}
	}

	return dp[len(s)]
}
