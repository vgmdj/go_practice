package Find_Minimum_in_Rotated_Sorted_Array

import "math"

func findMin(nums []int) int {
	var left, right, mid = 0, len(nums) - 1, 0

	for left <= right {
		mid = left + (right-left)/2
		if nums[left] <= nums[right] {
			return nums[left]
		}

		if nums[mid] < nums[right] {
			right = mid

		} else {
			left = mid + 1

		}

	}

	return math.MinInt32

}
