package Find_Minimum_in_Rotated_Sorted_Array_II

func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[right] {
			left = mid + 1

		} else if nums[mid] < nums[right] {
			right = mid

		} else {
			right = right - 1
		}

	}

	return nums[left]

}
