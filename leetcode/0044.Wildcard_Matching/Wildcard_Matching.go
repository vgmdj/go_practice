package Wildcard_Matching

import "strings"

func isMatch(s string, p string) bool {
	if len(s) <= 0 {
		if len(strings.Trim(p, "*")) <= 0 {
			return true
		}
		return false
	}

	dp := make([][]bool, len(p)+1)
	for i := 0; i < len(p)+1; i++ {
		dp[i] = make([]bool, len(s)+1)
	}

	dp[0][0] = true

	for i := 1; i < len(p)+1; i++ {
		if dp[i-1][0] == true && p[i-1] == '*' {
			dp[i][0] = true

		}
	}

	for i := 1; i < len(p)+1; i++ {
		for j := 1; j < len(s)+1; j++ {
			switch p[i-1] {
			case '*':
				dp[i][j] = dp[i-1][j] || dp[i][j-1]

			case '?':
				dp[i][j] = dp[i-1][j-1]

			default:
				dp[i][j] = dp[i-1][j-1] && p[i-1] == s[j-1]

			}

		}
	}

	return dp[len(p)][len(s)]
}
