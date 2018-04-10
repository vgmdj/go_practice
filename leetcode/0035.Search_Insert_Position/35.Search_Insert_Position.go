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

func searchInsertII(nums []int, target int) int {
	min := 0
	max := len(nums) - 1

	if len(nums) == 0 {
		return 0
	} else if nums[max] < target {
		return len(nums)
	} else if nums[0] > target {
		return 0
	}

	for {
		currentIndex := (max + min) / 2

		if nums[currentIndex] == target || nums[currentIndex] > target && nums[currentIndex-1] < target {
			return currentIndex
		} else if nums[currentIndex] > target {
			max = currentIndex - 1
		} else {
			min = currentIndex + 1
		}
	}

	return 0
}
