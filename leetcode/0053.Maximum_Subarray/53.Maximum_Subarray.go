package Maximum_Subarray

import "math"

// the most simple way
func maxSubArray(nums []int) int {
	max := 0
	if len(nums) > 0 {
		max = nums[0]
	}

	for i := 0; i < len(nums); i++ {
		result := nums[i]

		if result > max {
			max = result
		}

		for j := i + 1; j < len(nums); j++ {
			result += nums[j]
			if result > max {
				max = result
			}

		}

	}

	return max
}

// dp solution and O(n) space complexity
func maxSubArray2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums))
	dp[0] = nums[0]

	max := math.MinInt32

	for i := 1; i < len(nums); i++ {
		if nums[i]+dp[i-1] > nums[i] {
			dp[i] = nums[i] + dp[i-1]

		} else {
			dp[i] = nums[i]

		}

		if max < dp[i] {
			max = dp[i]
		}
	}

	return max
}

// dp solution and O(1) space complexity
func maxSubArray3(nums []int) int {
	var sum = math.MinInt32
	var result = math.MinInt32

	for i := 0; i < len(nums); i++ {
		sum = max(nums[i], sum+nums[i])
		result = max(result, sum)
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}




