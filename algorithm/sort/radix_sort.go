package sort

import "math"

func RadixSort(nums []int) {
	if len(nums) < 2 {
		return
	}

	maxValue := max(nums...)
	buf := make([]int, len(nums))
	for base := 1; base <= maxValue; base *= 10 {
		cnt := [10]int{}
		for _, v := range nums {
			c := v / base % 10
			cnt[c]++
		}

		for i := 1; i < len(cnt); i++ {
			cnt[i] += cnt[i-1]
		}

		for i := len(nums) - 1; i >= 0; i-- {
			c := nums[i] / base % 10
			buf[cnt[c]-1] = nums[i]
			cnt[c]--
		}

		copy(nums, buf)
	}

}

func max(nums ...int) int {
	if len(nums) == 0 {
		return math.MinInt32
	}

	maxValue := nums[0]
	for _, v := range nums {
		if v > maxValue {
			maxValue = v

		}
	}

	return maxValue
}
