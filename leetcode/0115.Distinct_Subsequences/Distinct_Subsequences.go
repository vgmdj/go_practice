package Distinct_Subsequences

func numDistinct(s string, t string) int {
	m, n := len(s), len(t)
	dp := make([][]int, m+1)
	dp[0] = make([]int, n+1)
	dp[0][0] = 1

	for i := 1; i <= m; i++ {
		dp[i] = make([]int, n+1)
		dp[i][0] = 1

		for j := 1; j <= n && j <= i; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + dp[i-1][j]

			} else {
				dp[i][j] = dp[i-1][j]

			}
		}
	}

	return dp[m][n]
}
