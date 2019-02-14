package _4Sum

import "sort"

func fourSum(nums []int, target int) [][]int {
	var result [][]int
	if len(nums) < 4 {
		return result
	}

	sort.Ints(nums)

	for first := 0; first < len(nums)-3; first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}

		for second := first + 1; second < len(nums)-2; second++ {
			if (second > first+1 && nums[second] == nums[second-1]) ||
				(nums[second]+nums[len(nums)-1]+nums[len(nums)-2] < target-nums[first]) {
				continue
			}

			third, forth := second+1, len(nums)-1

			for third < forth {
				if third > second+1 && nums[third] == nums[third-1] {
					third++
					continue
				}

				if forth < len(nums)-1 && nums[forth] == nums[forth+1] {
					forth--
					continue
				}

				sum := nums[second] + nums[third] + nums[forth]
				if sum == target-nums[first] {
					result = append(result, []int{nums[first], nums[second], nums[third], nums[forth]})
					third++

				} else if sum < target-nums[first] {
					third++

				} else {
					forth--
				}

			}

		}

	}

	return result

}
