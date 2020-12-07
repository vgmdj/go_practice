package Create_Maximum_Number

func maxNumber(nums1 []int, nums2 []int, k int) []int {
	start := 0
	if k > len(nums2) {
		start = k - len(nums2)
	}

	result := make([]int, 0)
	for i := start; i <= k && i <= len(nums1); i++ {
		s1 := maxSubSequence(nums1, i)
		s2 := maxSubSequence(nums2, k-i)
		mergedList := merge(s1, s2)
		if numberListLess(result, mergedList) {
			result = mergedList
		}
	}

	return result

}

func numberListLess(nums1, nums2 []int) bool {
	for i := 0; i < len(nums1) && i < len(nums2); i++ {
		if nums1[i] != nums2[i] {
			return nums1[i] < nums2[i]
		}
	}

	return len(nums1) < len(nums2)
}

func merge(nums1, nums2 []int) []int {
	result := make([]int, 0, len(nums1)+len(nums2))
	index1, index2 := 0, 0
	for i := 0; index1 < len(nums1) && index2 < len(nums2); i++ {
		if numberListLess(nums1[index1:], nums2[index2:]) {
			result = append(result, nums2[index2])
			index2++

		} else {
			result = append(result, nums1[index1])
			index1++

		}
	}

	if index1 < len(nums1) {
		result = append(result, nums1[index1:]...)
	}

	if index2 < len(nums2) {
		result = append(result, nums2[index2:]...)
	}

	return result

}

func maxSubSequence(nums []int, k int) []int {
	result := make([]int, 0)
	for i := 0; i < len(nums) && k > 0; i++ {
		for len(result) > 0 && len(nums)+len(result)-i-1 >= k && nums[i] > result[len(result)-1] {
			result = result[:len(result)-1]
		}

		if len(result) < k {
			result = append(result, nums[i])
		}

	}

	return result
}
