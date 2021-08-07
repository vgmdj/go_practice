package Pow_x_n_

import (
	"math"
)

func myPow(x float64, n int) float64 {
	if n > 0 {
		return quickPow(x, n)
	}

	return 1.0 / quickPow(x, -n)

}

func quickPow(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}

	y := quickPow(x, n/2)
	if n%2 == 0 {
		return y * y
	}

	return y * y * x

}
