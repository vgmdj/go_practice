package Median_of_Two_Sorted_Arrays

import (
	"sort"
)

// O(log(m+n))
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	//ensure n >= m
	m, n := len(nums1), len(nums2)
	if m > n {
		nums1, nums2 = nums2, nums1
		m, n = n, m
	}

	iMin, iMax, halfLen := 0, m, (m+n+1)/2
	for iMin <= iMax {
		i := (iMin + iMax) / 2
		j := halfLen - i
		if i < iMax && nums2[j-1] > nums1[i] {
			iMin = i + 1 // i is too small

		} else if i > iMin && nums1[i-1] > nums2[j] {
			iMax = i - 1 // i is too big

		} else { // i is perfect
			maxLeft := 0
			if i == 0 {
				maxLeft = nums2[j-1]

			} else if j == 0 {
				maxLeft = nums1[i-1]

			} else {
				maxLeft = Max(nums1[i-1], nums2[j-1])

			}

			if (m+n)%2 == 1 {
				return float64(maxLeft)
			}

			minRight := 0
			if i == m {
				minRight = nums2[j]

			} else if j == n {
				minRight = nums1[i]

			} else {
				minRight = Min(nums2[j], nums1[i])
			}

			return float64(maxLeft+minRight) / 2.0
		}
	}
	return 0.0
}

func Min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

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
