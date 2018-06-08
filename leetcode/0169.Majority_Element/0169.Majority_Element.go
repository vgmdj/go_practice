package Majority_Element

import "sort"

func majorityElement(nums []int) int {
	count := 0
	var data int
	for i := 0; i < len(nums); i++ {
		if count == 0 {
			data = nums[i]
		}

		if nums[i] == data {
			count += 1
		} else {
			count += -1
		}
	}
	return data
}

func majorityElement1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	sort.Ints(nums)

	return nums[len(nums)/2]

}
