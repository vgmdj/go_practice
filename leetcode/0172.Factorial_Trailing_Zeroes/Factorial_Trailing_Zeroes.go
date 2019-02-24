package Factorial_Trailing_Zeroes

func trailingZeroes(n int) int {
	var sum int
	for n != 0 {
		n /= 5
		sum += n
	}

	return sum

}
