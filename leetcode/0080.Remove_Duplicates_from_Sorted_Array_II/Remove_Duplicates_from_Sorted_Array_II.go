package Remove_Duplicates_from_Sorted_Array_II

func removeDuplicates(nums []int) int {
	var pre, cur = 1, 1

	for cur < len(nums) {
		if nums[cur] == nums[pre] && nums[pre] == nums[pre-1] {
			cur++
		} else if pre+1 == cur {
			pre++
			cur++

		} else if pre < len(nums)-1 {
			pre++
			nums[pre] = nums[cur]
			cur++
		}

	}

	return pre + 1
}
