package Median_of_Two_Sorted_Arrays

/*
m
n

i = (1+m)/2
i+j = (n+m)/2
j=(n+m)/2-i
j>=0
n+m >= 2m
n>=m



left      |     right
1,2,3...i | i+1,.....m
1,2,3...j | j+1,.....n


*/

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
