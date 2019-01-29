package House_Robber

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums))
	if len(nums) >= 1 {
		dp[0] = nums[0]
	}

	if len(nums) >= 2 {
		dp[1] = Max(nums[1], dp[0])
	}

	for i := 2; i < len(nums); i++ {
		dp[i] = Max(dp[i-1], dp[i-2]+nums[i])
	}

	return dp[len(nums)-1]
}

func Max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
