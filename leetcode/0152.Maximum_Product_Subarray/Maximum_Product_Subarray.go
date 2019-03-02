package Maximum_Product_Subarray

import "math"

func maxProduct(nums []int) int {
	if len(nums) < 1 {
		return 0
	}

	var result = math.MinInt32
	min, max := 0, 0
	for k, v := range nums {
		if k == 0 {
			max, min = v, v
			result = v
			continue
		}
		max, min = compare(v, v*min, v*max)
		result, _ = compare(result, max)
	}

	return result

}

func compare(a ...int) (max, min int) {
	if len(a) < 1 {
		return math.MinInt32, math.MaxInt32
	}

	max, min = a[0], a[0]
	for i := 1; i < len(a); i++ {
		if a[i] > max {
			max = a[i]
		}

		if a[i] < min {
			min = a[i]
		}

	}

	return

}
