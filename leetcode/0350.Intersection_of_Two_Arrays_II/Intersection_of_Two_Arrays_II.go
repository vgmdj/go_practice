package Intersection_of_Two_Arrays_II

func intersect(nums1 []int, nums2 []int) []int {
	same := make(map[int]int, len(nums1)+len(nums2))
	var result []int

	for _, v := range nums1 {
		same[v]++
	}

	for _, v := range nums2 {
		if same[v] != 0 {
			result = append(result, v)
			same[v]--
		}

	}

	return result

}
