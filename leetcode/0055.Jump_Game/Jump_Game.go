package Jump_Game

func canJump(nums []int) bool {
	max := 0
	for i := 0; i <= max; i++ {
		if i+nums[i] > max {
			max = i + nums[i]
		}
		if max >= len(nums)-1 {
			return true
		}
	}
	return false
}

func canJump2(nums []int) bool {
	p := 0
	needJump := false

	for i := len(nums) - 2; i >= 0; i-- {
		if needJump {
			if nums[i] < p-i+1 {
				needJump = true
				continue
			}

			needJump = false
			p = i
		}

		if nums[i] == 0 {
			p = i
			needJump = true
		}
	}

	return !needJump

}