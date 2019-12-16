package Reverse_Integer

import "math"

func reverse(x int) int {
	if x > -10 && x < 10 {
		return x
	}

	result := 0
	for x != 0 {
		c := x % 10
		x /= 10
		if result > math.MaxInt32/10 || (result == math.MaxInt32/10 && c > math.MaxInt32%10) ||
			result < math.MinInt32/10 || (result == math.MinInt32/10 && c < math.MinInt32%10) {
			return 0
		}

		result = result*10 + c
	}

	return result
}
