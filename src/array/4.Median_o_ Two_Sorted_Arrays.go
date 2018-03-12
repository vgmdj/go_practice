package main

import (
	"log"
	"sort"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	nums1 = append(nums1, nums2...)

	sort.Ints(nums1)

	length := len(nums1)

	if length%2 == 1 {
		return float64(nums1[length/2])
	}

	return float64(nums1[length/2]+nums1[length/2-1]) / 2.0

}

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}

	log.Println(findMedianSortedArrays(nums1, nums2))

}
