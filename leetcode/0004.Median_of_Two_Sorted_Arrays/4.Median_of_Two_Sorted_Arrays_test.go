package Median_of_Two_Sorted_Arrays

import "testing"

func TestOK(t *testing.T) {
	nums1 := []int{1, 3, 5, 7, 9}
	nums2 := []int{2, 3, 4, 5, 6, 7, 8, 9}

	t.Log(findMedianSortedArrays(nums1, nums2))
}
