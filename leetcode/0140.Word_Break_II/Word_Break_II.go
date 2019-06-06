package Word_Break_II

func wordBreak(s string, wordDict []string) []string {
	result := make([]string, 0)
	possibility := make([]int, len(s))
	for i := range possibility {
		possibility[i] = -1
	}

	helper(s, wordDict, 0, possibility, "", &result)

	return result

}

func helper(s string, wordDict []string, index int, possibility []int, subString string, result *[]string) int {
	if index >= len(s) {
		*result = append(*result, subString[1:])
		return 1
	}

	if possibility[index] == 0 {
		return 0
	}

	sum := 0
	for _, word := range wordDict {
		if index+len(word) > len(s) || s[index:index+len(word)] != word {
			continue
		}

		sum = sum + helper(s, wordDict, index+len(word), possibility, subString+" "+word, result)

	}

	possibility[index] = sum

	return sum

}
