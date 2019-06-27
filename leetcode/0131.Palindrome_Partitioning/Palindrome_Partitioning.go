package Palindrome_Partitioning

func partition(s string) [][]string {
	if len(s) == 0 {
		return [][]string{}
	}

	results := make([][]string, 0)

	backTracking(s, &results, []string{}, 0)

	return results
}

func backTracking(s string, result *[][]string, subset []string, index int) {
	if index == len(s) {
		*result = append(*result, append([]string{}, subset...))
		return
	}

	for i := 1; index+i <= len(s); i++ {
		if !isPalindrome(s[index : index+i]) {
			continue
		}

		backTracking(s, result, append(subset, s[index:index+i]), index+i)

	}

}

func partition2(s string) [][]string {
	if len(s) == 1 {
		return [][]string{{s}}
	}

	results := make([][]string, 0)

	for i := 0; i < len(s); i++ {
		if !isPalindrome(s[:i+1]) {
			continue
		}

		before := s[:i+1]

		after := partition(s[i+1:])
		if len(after) == 0 {
			results = append(results, []string{before})
			continue
		}

		for _, v := range after {
			result := append([]string{}, before)
			result = append(result, v...)

			results = append(results, result)
		}

	}

	return results
}

func isPalindrome(subset string) bool {
	if len(subset) == 0 {
		return false
	}

	if len(subset) == 1 {
		return true
	}

	left, right := 0, len(subset)-1
	for left <= right {
		if subset[left] != subset[right] {
			return false
		}

		left++
		right--

	}

	return true

}
