package Edit_Distance

func minDistance(word1 string, word2 string) int {
	//存在某一个为空
	if len(word1) == 0 || len(word2) == 0 {
		return len(word1) + len(word2)
	}

	//初始化 dp
	dp := make([][]int, len(word1)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(word2)+1)
	}

	//第一列
	for i := 0; i < len(dp); i++ {
		dp[i][0] = i
	}

	//第一行
	for i := 0; i < len(dp[0]); i++ {
		dp[0][i] = i
	}

	for w1 := 1; w1 <= len(word1); w1++ {
		for w2 := 1; w2 <= len(word2); w2++ {
			if word2[w2-1] == word1[w1-1] {
				dp[w1][w2] = dp[w1-1][w2-1]
				continue
			}

			dp[w1][w2] = 1 + Min(dp[w1-1][w2-1], dp[w1-1][w2], dp[w1][w2-1])

		}

	}

	return dp[len(word1)][len(word2)]
}

func Min(a, b, c int) int {
	min := a
	if b < min {
		min = b
	}

	if c < min {
		min = c
	}

	return min

}
