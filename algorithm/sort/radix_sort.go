package sort

import "math"

func RadixSort(nums []int) {
	base, length := 1, getLen(nums)

	for i := 1; i <= length; i++ {
		radixSort(nums, base)
		base *= 10
	}

}

func getLen(nums []int) int {
	max, length := math.MinInt32, 0

	for _, v := range nums {
		if v > max {
			max = v
		}
	}

	for max != 0 {
		max /= 10
		length++
	}

	return length
}

func radixSort(nums []int, base int) {
	bucket := make([][]int, 10)

	for _, v := range nums {
		k := v / base % 10
		bucket[k] = append(bucket[k], v)

	}

	index := 0
	for _, group := range bucket {
		for _, v := range group {
			nums[index] = v
			index++
		}
	}

}
