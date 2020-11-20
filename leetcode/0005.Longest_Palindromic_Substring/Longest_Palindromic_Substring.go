package Longest_Palindromic_Substring

func longestPalindrome(s string) string {
	var result string
	maxLength := 0
	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}

	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			if i != j && (s[i] != s[j] || (j-i != 1 && !dp[i+1][j-1])) {
				continue
			}

			dp[i][j] = true

			if j-i+1 > maxLength {
				maxLength = j - i + 1
				result = s[i : j+1]
			}

		}
	}

	return result
}

func longestPalindrome2(s string) string {
	if len(s) <= 1 {
		return s
	}

	var start, length, longest int
	for i := 0; i < len(s); i++ {
		length = findLongestPalindrome(s, i, i)
		if length > longest {
			longest = length
			start = i - length/2
		}

		length = findLongestPalindrome(s, i, i+1)
		if length > longest {
			longest = length
			start = i - length/2 + 1
		}
	}

	return s[start : start+longest]

}

func findLongestPalindrome(s string, left, right int) int {
	var length int
	for left >= 0 && right < len(s) {
		if s[left] != s[right] {
			break
		}

		length = right - left + 1
		left--
		right++
	}
	return length
}
