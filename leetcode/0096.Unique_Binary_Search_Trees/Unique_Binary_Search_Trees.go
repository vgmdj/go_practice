package Unique_Binary_Search_Trees

func numTrees(n int) int {
	if n < 3 {
		return n
	}

	var dp = make([]int, n+1)
	dp[0], dp[1], dp[2] = 1, 1, 2

	for i := 3; i <= n; i++ {
		for j := 0; j < i; j++ {
			dp[i] += dp[j] * dp[i-j-1]
		}
	}

	return dp[n]

}
