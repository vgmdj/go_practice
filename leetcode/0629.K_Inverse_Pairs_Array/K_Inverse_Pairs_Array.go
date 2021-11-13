package K_Inverse_Pairs_Array

func kInversePairs(n int, k int) int {
	if k == 0 {
		return 1
	}
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, k+1)
	}

	for i := 1; i < n+1; i++ {
		dp[i][0] = 1

		for j := 1; j < k+1; j++ {
			for m := 0; m <= j; m++ {
				dp[i][j] += dp[i-1][j-m]
			}
		}
	}

	return dp[n][k]

}
