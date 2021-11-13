package Guess_Number_Higher_or_Lower_II

import "math"

func getMoneyAmount(n int) int {
	dp := make([][]int, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+1)
	}

	for i := n-1; i > 0; i-- {
		for j := i + 1; j <= n; j++ {
			minCost := math.MaxInt32
			for k := i; k < j; k++ {
				cost := k + max(dp[i][k-1], dp[k+1][j])
				minCost = min(minCost, cost)
			}
			dp[i][j] = minCost
		}
	}

	return dp[1][n]

}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
