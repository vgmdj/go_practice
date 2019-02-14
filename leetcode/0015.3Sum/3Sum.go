package _3Sum

import "sort"

func threeSum(nums []int) [][]int {
	var result [][]int

	if len(nums) < 3 {
		return result
	}

	sort.Ints(nums)

	for first := 0; first < len(nums)-2 && nums[first] <= 0; first++ {
		if (first > 0 && nums[first] == nums[first-1]) || (nums[first]+nums[len(nums)-1]+nums[len(nums)-2] < 0) {
			continue
		}

		second, third := first+1, len(nums)-1

		for second < third {
			if second > first+1 && nums[second] == nums[second-1] {
				second++
				continue
			}

			if third < len(nums)-1 && nums[third] == nums[third+1] {
				third--
				continue
			}

			sum := nums[first] + nums[second] + nums[third]
			if sum == 0 {
				result = append(result, []int{nums[first], nums[second], nums[third]})
				second++

			} else if sum < 0 {
				second++

			} else {
				third--

			}

		}

	}

	return result

}
