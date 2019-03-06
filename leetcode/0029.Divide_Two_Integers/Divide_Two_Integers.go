package Divide_Two_Integers

import "math"

func divide(dividend int, divisor int) int {
	quotient := 0
	sign := 1

	//the abs minInt32 is equal maxInt32+1, so if the divisor = -1, will overflow
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	// capture the sign
	if divisor < 0 && dividend < 0 {
		divisor = -divisor
		dividend = -dividend

	} else if divisor < 0 {
		divisor = -divisor
		sign = -1

	} else if dividend < 0 {
		dividend = -dividend
		sign = -1

	}

	// divide some part to approach the answer
	for dividend >= divisor {
		multiple := divisor
		for i := 0; dividend >= multiple; i++ {
			dividend -= multiple
			quotient += 1 << uint(i)
			multiple <<= 1
		}
	}

	if sign == -1 {
		return -quotient
	}

	return quotient
}
