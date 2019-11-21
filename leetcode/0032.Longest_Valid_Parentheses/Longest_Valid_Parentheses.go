package Longest_Valid_Parentheses

func longestValidParentheses(s string) int {
	max := 0
	dp := make([]int, len(s))

	for i := 1; i < len(s); i++ {
		if s[i] != ')' {
			continue
		}

		if s[i-1] == '(' {
			if i >= 2 {
				dp[i] = dp[i-2] + 2

			} else {
				dp[i] = 2

			}

		} else if i-dp[i-1] > 0 && s[i-dp[i-1]-1] == '(' {
			dp[i] = dp[i-1] + 2

			if i-dp[i-1] >= 2 {
				dp[i] += dp[i-dp[i-1]-2]
			}
		}

		max = Max(max, dp[i])

	}

	return max

}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestValidParentheses2(s string) int {
	counter := 0
	record := make([]bool, len(s))
	for k, v := range s {
		if v != '(' && v != ')' {
			return 0
		}

		if v == '(' {
			counter++
			continue

		}

		if v == ')' {
			if counter <= 0 {
				continue
			}

			counter--
			record[k] = true
		}
	}

	counter = 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != '(' && s[i] != ')' {
			return 0
		}

		if s[i] == ')' {
			counter++
			continue

		}

		if s[i] == '(' {
			if counter <= 0 {
				continue
			}

			counter--
			record[i] = true
		}
	}

	max, c := 0, 0
	for i := 0; i < len(record); i++ {
		if !record[i] {
			c = 0
			continue
		}

		if c == 0 || (c != 0 && record[i-1]) {
			c++

		}

		if max < c {
			max = c
		}

	}

	return max

}
