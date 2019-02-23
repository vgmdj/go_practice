package Search_in_Rotated_Sorted_Array_II

func search(nums []int, target int) bool {
	var left, right, mid = 0, len(nums) - 1, 0

	for left <= right {
		mid = left + (right-left)/2
		if nums[mid] == target {
			return true
		}

		if nums[left] == nums[mid] {
			left++

		} else if nums[left] < nums[mid] {
			if nums[left] <= target && target <= nums[mid] {
				right = mid - 1

			} else {
				left = mid + 1
			}

		} else {
			if nums[mid] <= target && target <= nums[right] {
				left = mid + 1

			} else {
				right = mid - 1

			}
		}
	}

	return false

}
