package Intersection_of_Two_Arrays

import "sort"

func intersection(nums1 []int, nums2 []int) []int {
	var result []int
	var isSame = make(map[int]bool)

	for _, v := range nums1 {
		isSame[v] = true
	}

	for k, v := range nums2 {
		if isSame[v] {
			result = append(result, nums2[k])
			isSame[v] = false
		}
	}

	return result
}

func intersection2(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)

	var result []int
	var p, q = 0, 0
	for p < len(nums1) && q < len(nums2) {
		if p > 0 && nums1[p] == nums1[p-1] {
			p++
			continue
		}

		if q > 0 && nums2[q] == nums2[q-1] {
			q++
			continue
		}

		if nums1[p] == nums2[q] {
			result = append(result, nums1[p])
			p++
			q++
			continue
		}

		if nums1[p] < nums2[q] {
			p++
			continue
		}

		q++

	}

	return result

}
