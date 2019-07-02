package Best_Time_to_Buy_and_Sell_Stock_with_Cooldown

/*
对于天数来说，每天都有两个状态，即持有，不持有

- 持有状态，可由持有和两天前的不持有状态转移而来
- 不持有状态，可由不持有和持有状态转移而来

而两天前的不持有转持有，则是买入，需使 profit-price[i]
一天前的持有转不持有，则是卖出，需使 profit+price[i]


最后一天的最大利润，一定是在不持有里，即结果是 dp[len(price)-1][0]

状态转移方程:

dp[i][0] = Max( dp[i-1][0], dp[i-1][1]+price[i] )
dp[i][1] = Max( dp[i-1][1], dp[i-2][0]-price[i] )

base case:

dp[0][0] = 0
dp[0][1] = -price[0]

*/

func maxProfit1(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	dp := make([][2]int, len(prices))
	dp[0][0] = 0
	dp[0][1] = -prices[0]

	for i := 1; i < len(prices); i++ {
		dp[i][0] = Max(dp[i-1][0], dp[i-1][1]+prices[i])

		if i == 1 {
			dp[i][1] = Max(dp[i-1][1], -prices[i])

		} else {
			dp[i][1] = Max(dp[i-1][1], dp[i-2][0]-prices[i])

		}

	}

	return dp[len(prices)-1][0]

}

// simplify
func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	dp_i1, dp_i0, dp_pre := -prices[0], 0, 0

	for i := 1; i < len(prices); i++ {
		tmp := dp_i0
		dp_i0 = Max(dp_i0, dp_i1+prices[i])
		dp_i1 = Max(dp_i1, dp_pre-prices[i])
		dp_pre = tmp

	}

	return dp_i0

}

func Max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
