package Unique_Paths

func uniquePaths(m int, n int) int {
	if n <= 1 || m <= 1 {
		return 1
	}

	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < len(dp); j++ {
			dp[j] = dp[j] + dp[j-1]
		}
	}

	return dp[n-1]
}