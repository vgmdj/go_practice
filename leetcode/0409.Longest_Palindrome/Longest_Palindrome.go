package Longest_Palindrome

func longestPalindrome(s string) int {
	hasOdd := false
	counter := make([]int, 128)
	maxLength := 0

	for _, v := range s {
		counter[v]++
	}

	for _, v := range counter {
		if v%2 == 0 {
			maxLength += v

		} else {
			hasOdd = true

			if v > 1 {
				maxLength += v - 1
			}

		}
	}

	if hasOdd {
		return maxLength + 1
	}

	return maxLength

}
