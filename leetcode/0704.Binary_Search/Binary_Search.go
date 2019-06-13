package Binary_Search

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}

		if nums[mid] < target {
			left = mid + 1

		} else {
			right = mid - 1

		}

	}

	return -1

}

func search2(nums []int, target int) int {
	return helper(nums, target, 0, len(nums)-1)
}

func helper(nums []int, target, start, end int) int {
	if len(nums) < 1 || start < 0 || end > len(nums)-1 || end < start {
		return -1
	}

	mid := (start + end) / 2

	if nums[mid] == target {
		return mid
	}

	if nums[mid] < target {
		return helper(nums, target, mid+1, end)
	}

	return helper(nums, target, start, mid-1)

}
