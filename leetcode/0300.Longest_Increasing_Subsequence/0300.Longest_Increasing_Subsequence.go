package Longest_Increasing_Subsequence

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	max := 0
	dp := make([]int, len(nums))
	for i := 1; i < len(nums); i++ {
		for j := i - 1; j >= 0; j-- {
			if nums[i] > nums[j] {
				dp[i] = Max(dp[j]+1, dp[i])
			}
		}

		max = Max(max,dp[i])

	}

	return max+1
}

func Max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
