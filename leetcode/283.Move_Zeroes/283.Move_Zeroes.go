package Move_Zeroes

func moveZeroes(nums []int) {
	if len(nums) <= 1 {
		return
	}

	lastNonZeroFoundAt, cur := 0, 0
	for ; cur < len(nums); cur++ {
		if nums[cur] != 0 {
			nums[lastNonZeroFoundAt] = nums[cur]
			lastNonZeroFoundAt++
		}
	}

	for zero := lastNonZeroFoundAt; zero < len(nums); zero++ {
		nums[zero] = 0
	}

}

//the better
func moveZeroesII(nums []int) {
	if len(nums) <= 1 {
		return
	}

	for lastNonZeroFoundAt, cur := 0, 0; cur < len(nums); cur++ {
		if nums[cur] != 0 {
			nums[lastNonZeroFoundAt], nums[cur] = nums[cur], nums[lastNonZeroFoundAt]
			lastNonZeroFoundAt++
		}
	}

}
