package Product_of_Array_Except_Self

func productExceptSelf(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	var result = make([]int, len(nums))
	var left = 1
	var right = 1

	for i := 0; i < len(nums); i++ {
		result[i] = left
		left *= nums[i]
	}

	for i := len(nums) - 1; i >= 0; i-- {
		result[i] *= right
		right *= nums[i]
	}

	return result

}
