package Wiggle_Subsequence

func wiggleMaxLength(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	var max, cur, flag = 1, 1, 0

	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] > 0 && flag <= 0 {
			cur++
			flag = 1
			max = Max(max, cur)
			continue

		}

		if nums[i]-nums[i-1] < 0 && flag >= 0 {
			cur++
			flag = -1
			max = Max(max, cur)
			continue

		}

	}

	return max
}

func Max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
