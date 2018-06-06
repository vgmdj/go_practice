package SqrtX

import "math"

func mySqrt(x int) int {
	if x <= 0 {
		return 0
	}

	return int(math.Sqrt(float64(x)))
}
