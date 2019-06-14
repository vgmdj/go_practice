package Minimum_Size_Subarray_Sum

import "math"

func minSubArrayLen(s int, nums []int) int {
	sum, count, left, right := 0, math.MaxInt32, 0, 0

	for ; right < len(nums); right++ {
		sum += nums[right]

		for sum >= s {
			count = Min(count, right-left+1)
			sum -= nums[left]
			left++
		}

	}

	if count == math.MaxInt32 {
		return 0
	}

	return count

}

func Min(a, b int) int {
	if a < b {
		return a
	}

	return b

}
