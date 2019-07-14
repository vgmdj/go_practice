package Integer_Break

func integerBreak(n int) int {
	if n == 2 || n == 3 {
		return n - 1
	}

	result := 1
	for n > 4 {
		n -= 3
		result *= 3

	}

	return result * n

}
