package main

import (
	"log"
	"sort"
)

// O((m+n)log(m+n))
func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {

	nums1 = append(nums1, nums2...)

	sort.Ints(nums1)

	length := len(nums1)

	if length%2 == 1 {
		return float64(nums1[length/2])
	}

	return float64(nums1[length/2]+nums1[length/2-1]) / 2.0

}

// O(m+n)
func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	var nums3 []int
	var midl, midr int
	var left1, left2 int

	if (len(nums1)+len(nums2))%2 == 1 {
		midl = (len(nums1) + len(nums2)) / 2
		midr = midl
	} else {
		midl = (len(nums1)+len(nums2))/2 - 1
		midr = midl + 1
	}

	for left1 < len(nums1) && left2 < len(nums2) {
		if nums1[left1] < nums2[left2] {
			nums3 = append(nums3, nums1[left1])
			left1++
			continue
		}
		nums3 = append(nums3, nums2[left2])
		left2++

	}

	nums3 = append(nums3, nums1[left1:]...)
	nums3 = append(nums3, nums2[left2:]...)

	return float64(nums3[midl]+nums3[midr]) / 2.0

}

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}

	log.Println(findMedianSortedArrays2(nums1, nums2))

}
