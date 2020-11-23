package Summary_Ranges

import "strconv"

func summaryRanges(nums []int) []string {
	result := make([]string, 0)
	for i, j := 0, 0; j < len(nums); j++ {
		if j+1 < len(nums) && nums[j]+1 == nums[j+1] {
			continue
		}

		if i == j {
			result = append(result, strconv.Itoa(nums[i]))

		} else {
			result = append(result, strconv.Itoa(nums[i])+"->"+strconv.Itoa(nums[j]))

		}
		i = j + 1

	}

	return result

}
