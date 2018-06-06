package Median_of_Two_Sorted_Arrays

import "testing"

func TestOK(t *testing.T) {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}

	t.Log(findMedianSortedArrays2(nums1, nums2))
}
