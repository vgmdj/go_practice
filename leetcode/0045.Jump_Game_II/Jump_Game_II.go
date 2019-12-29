package Jump_Game_II

func jump(nums []int) int {
	if len(nums) <= 1 || nums[0] == 0 {
		return 0
	}

	max, result, reach := nums[0], 0, nums[0]
	for i := 0; i < len(nums); i++ {
		if i+nums[i] > max {
			max = i + nums[i]
		}

		if i == reach {
			result++
			reach = max
		}

		if reach >= len(nums)-1 {
			return result +1
		}

	}

	return result

}
