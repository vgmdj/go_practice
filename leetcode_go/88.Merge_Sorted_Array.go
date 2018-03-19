package main

import "log"

func merge(nums1 []int, m int, nums2 []int, n int) {
	for p := m + n - 1; p >= 0; p-- {
		if m == 0 && n != 0 {
			nums1[p] = nums2[n-1]
			n--
		} else if n == 0 && m != 0 {
			nums1[p] = nums1[m-1]
			m--
		} else if nums1[m-1] < nums2[n-1] {
			nums1[p] = nums2[n-1]
			n--
		} else {
			nums1[p] = nums1[m-1]
			m--
		}

	}

}

func main() {
	nums1 := []int{1, 0, 0}
	nums2 := []int{1, 2}
	merge(nums1, 0, nums2, 2)

	log.Println(nums1)
}
