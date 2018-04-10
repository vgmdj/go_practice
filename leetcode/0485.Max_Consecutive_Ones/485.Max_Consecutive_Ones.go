package Max_Consecutive_Ones

func findMaxConsecutiveOnes(nums []int) int {

	count, max := 0, 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			count++
		} else {
			count = 0
		}

		if max < count {
			max = count
		}

	}

	return max

}
