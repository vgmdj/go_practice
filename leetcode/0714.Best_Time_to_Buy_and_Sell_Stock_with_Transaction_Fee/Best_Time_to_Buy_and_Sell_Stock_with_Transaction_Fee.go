package Best_Time_to_Buy_and_Sell_Stock_with_Transaction_Fee

func maxProfit(prices []int, fee int) int {
	if len(prices) == 0 {
		return 0
	}

	dp0, dp1 := 0, -prices[0]
	for i := 1; i < len(prices); i++ {
		dp0 = max(dp0, dp1+prices[i]-fee)
		dp1 = max(dp1, dp0-prices[i])

	}

	return max(dp0, dp1)

}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
