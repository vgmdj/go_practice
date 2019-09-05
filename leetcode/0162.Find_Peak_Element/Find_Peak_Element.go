package Find_Peak_Element

// O(lgN) non-recursion
func findPeakElement(nums []int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		if left == right {
			return left
		}

		mid := (left + right) / 2
		if nums[mid] > nums[mid+1] {
			right = mid

		} else {
			left = mid + 1

		}
	}

	return -1

}

// O(logN) recursion
func findPeakElement_Rec(nums []int) int {
	return search(nums, 0, len(nums)-1)

}

func search(nums []int, start, end int) int {
	if end < 0 || end < start {
		return -1
	}

	if start == end {
		return start
	}

	mid := (start + end) / 2
	if nums[mid] > nums[mid+1] {
		return search(nums, start, mid)
	}

	return search(nums, mid+1, end)

}

// O(n)
func findPeakElement_N(nums []int) int {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			return i
		}
	}

	return len(nums) - 1

}
