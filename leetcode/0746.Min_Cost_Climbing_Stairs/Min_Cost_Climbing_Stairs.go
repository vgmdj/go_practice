package Min_Cost_Climbing_Stairs

func minCostClimbingStairs(cost []int) int {
	dp, dp1, dp2 := 0, 0, 0

	for i := 2; i <= len(cost); i++ {
		dp = min(dp1+cost[i-1], dp2+cost[i-2])
		dp2 = dp1
		dp1 = dp
	}

	return dp

}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
