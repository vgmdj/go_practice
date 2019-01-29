package House_Robber_II

func rob(nums []int) int {
	if len(nums) < 4 {
		return Max(nums[:]...)
	}

	return Max(robPart(nums[1:]),robPart(nums[:len(nums)-1]))

}

func robPart(nums []int) int {
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

func Max(nums ...int) int {
	max := 0
	for _, v := range nums {
		if max < v {
			max = v
		}
	}

	return max
}
