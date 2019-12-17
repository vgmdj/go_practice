package Search_Insert_Position

func searchInsert(nums []int, target int) int {
	var position int
	for k, v := range nums {
		if v == target {
			return k
		}

		if v < target {
			position++
		}

	}

	return position

}

func searchInsert2(nums []int, target int) int {
	if len(nums) == 0 {
		return 0

	} else if nums[len(nums)-1] < target {
		return len(nums)

	} else if nums[0] > target {
		return 0

	}

	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}

		if nums[mid] > target {
			right = mid - 1

		} else {
			left = mid + 1

		}
	}

	return left
}
