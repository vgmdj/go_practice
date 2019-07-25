package Contains_Duplicate_III

// similar to bucket sort
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	if t < 0 || k <= 0 || len(nums) < 2 {
		return false
	}

	buckets := make(map[int]int, len(nums))

	for i := 0; i < len(nums); i++ {
		c := nums[i]
		if c < 0 {
			c -= t + 1
		}
		key := c / (t + 1)

		if _, ok := buckets[key]; ok {
			return true
		}

		if v, ok := buckets[key+1]; ok && Abs(v-nums[i]) < t+1 {
			return true
		}

		if v, ok := buckets[key-1]; ok && Abs(v-nums[i]) < t+1 {
			return true
		}

		buckets[key] = nums[i]

		if i-k >= 0 {
			d := nums[i-k]
			if d < 0 {
				d -= t + 1
			}

			delete(buckets, d/(t+1))
		}

	}

	return false
}

// sliding window
func containsNearbyAlmostDuplicate2(nums []int, k int, t int) bool {
	if k <= 0 || len(nums) < 2 {
		return false
	}

	left, right := 0, 1
	for left < right && right < len(nums) {

		for c := left; c < right; c++ {
			sub := nums[right] - nums[c]
			if Abs(sub) <= t {
				return true
			}
		}

		if right-left == k {
			right++
			left++

		} else if right-left < k {
			right++

		} else {
			left++

		}

	}

	return false

}

func Abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}
