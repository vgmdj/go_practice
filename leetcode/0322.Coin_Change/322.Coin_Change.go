package Coin_Change

import (
	"math"
)

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt32
	}
	for i := 1; i <= amount; i++ {
		for _, coinValue := range coins {
			if i-coinValue >= 0 {
				dp[i] = int(math.Min(float64(dp[i]), float64(dp[i-coinValue]+1)))
			}
		}
	}
	if dp[amount] == math.MaxInt32 {
		return -1
	}

	return dp[amount]
}
