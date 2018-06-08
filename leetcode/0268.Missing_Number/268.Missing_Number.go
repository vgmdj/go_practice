package Missing_Number

func missingNumber(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		if nums[0] == 0 {
			return 1
		}
		return 0
	}

	sum := 0

	for _, v := range nums {
		sum += v
	}

	return (len(nums)*(len(nums)+1))>>1 - sum

}
