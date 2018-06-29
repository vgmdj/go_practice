package Contains_Duplicate

import "sort"

func containsDuplicate(nums []int) bool {
	counter := make(map[int]struct{})

	for _, v := range nums {
		if _, ok := counter[v]; ok {
			return true
		} else {
			counter[v] = struct{}{}
		}
	}

	return false
}

func containsDuplicateII(nums []int) bool {
	sort.Ints(nums)

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return true
		}
	}

	return false
}
