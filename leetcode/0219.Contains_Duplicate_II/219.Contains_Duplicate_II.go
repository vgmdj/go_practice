package Contains_Duplicate_II

func containsNearbyDuplicate(nums []int, k int) bool {
	if len(nums) <= 1 || k < 1 {
		return false
	}

	counter := make(map[int]int, len(nums))

	for p, v := range nums {
		if lastP, ok := counter[v]; ok {
			if p-lastP <= k {
				return true
			}
		}
		counter[v] = p
	}

	return false
}
