package Longest_Increasing_Subsequence

func lengthOfLIS(nums []int) int {
	increasingArray := make([]int, 0)

	replaceFirstBigerOne := func(value int) {
		left, right := 0, len(increasingArray)-1
		for left <= right {
			mid := right - (right-left)/2
			if increasingArray[mid] == value {
				return
			}

			if increasingArray[mid] > value {
				right = mid - 1

			} else {
				left = mid + 1

			}

		}

		increasingArray[left] = value
	}

	for _, c := range nums {
		if len(increasingArray) == 0 {
			increasingArray = append(increasingArray, c)
			continue
		}

		left, right := increasingArray[0], increasingArray[len(increasingArray)-1]
		if c == left || c == right {
			continue
		}

		if c > right {
			increasingArray = append(increasingArray, c)
			continue
		}

		if c < left {
			increasingArray[0] = c
			continue
		}

		replaceFirstBigerOne(c)

	}

	return len(increasingArray)
}
