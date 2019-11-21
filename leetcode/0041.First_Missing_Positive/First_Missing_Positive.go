package First_Missing_Positive

func firstMissingPositive(nums []int) int {
	for i := 0; i < len(nums); {
		if nums[i] > 0 && nums[i] <= len(nums) && nums[nums[i]-1] != nums[i] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
			continue
		}

		i++
	}

	for k, v := range nums {
		if k+1 != v {
			return k + 1
		}
	}

	return len(nums) + 1

}

func firstMissingPositive1(nums []int) int {
	hashStr := make([]byte, len(nums)+2)
	for _, v := range nums {
		if v >= 0 && v <= len(nums) {
			hashStr[v]++
		}

	}

	for i := 1; i < len(hashStr); i++ {
		if hashStr[i] == 0 {
			return i
		}
	}

	return 1
}
