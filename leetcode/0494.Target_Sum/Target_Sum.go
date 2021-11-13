package Target_Sum

func findTargetSumWays(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, 2001)
	dp[nums[0]+1000] = 1
	dp[-nums[0]+1000] += 1

	for i := 1; i < len(nums); i++ {
		next := make([]int, 2001)
		for sum := -1000; sum < 1001; sum++ {
			if dp[sum+1000] > 0 {
				next[sum+nums[i]+1000] += dp[sum+1000]
				next[sum-nums[i]+1000] += dp[sum+1000]
			}
		}
		dp = next
	}

	return dp[target+1000]

}
