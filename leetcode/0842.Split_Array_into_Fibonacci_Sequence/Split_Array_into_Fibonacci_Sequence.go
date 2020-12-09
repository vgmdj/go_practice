package Split_Array_into_Fibonacci_Sequence

import "math"

func splitIntoFibonacci(s string) []int {
	result := make([]int, 0)
	backTracking(&result, s, []int{}, 0)
	return result
}

func backTracking(result *[]int, nums string, pred []int, lastTwoSum int) bool {
	c := 0
	for i := 0; i < len(nums); i++ {
		if (i != 0 && c == 0) || len(*result) > 0 {
			return false
		}

		c = c*10 + int(nums[i]-'0')
		if c > math.MaxInt32 {
			return false
		}

		if len(pred) < 2 {
			backTracking(result, nums[i+1:], append(pred, c), lastTwoSum+c)
			continue
		}

		if c > lastTwoSum {
			return false

		} else if c < lastTwoSum {
			continue

		}

		if i == len(nums)-1 {
			*result = append(pred, c)
			return true
		}

		return backTracking(result, nums[i+1:], append(pred, c), lastTwoSum+c-pred[len(pred)-2])

	}

	return false

}
