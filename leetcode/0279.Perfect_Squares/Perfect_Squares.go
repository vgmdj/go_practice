package _279_Perfect_Squares

import "math"

func numSquares(n int) int {
	if n < 4 {
		return n
	}

	var squares []int
	for i := 1; i <= n; i++ {
		square := i * i
		if square == n {
			return 1
		}

		if square > n {
			break
		}

		squares = append(squares, square)

	}

	var dp = make([]int, n+1)
	dp[0], dp[1] = 0, 1
	for i := 2; i <= n; i++ {
		dp[i] = math.MaxInt32
		for _, squ := range squares {
			if squ == i {
				dp[i] = 1
			} else if squ > i {
				break
			}

			dp[i] = Min(dp[i], 1+dp[i-squ])
		}
	}

	return dp[n]

}

func Min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
