package _397_Integer_Replacement

func integerReplacement(n int) int {
	if n <= 1 {
		return 0
	}

	if n%2 == 0 {
		n /= 2

	} else if (n+1)%4 != 0 || n == 3 {
		n = n - 1

	} else {
		n = n + 1

	}

	return integerReplacement(n) + 1

}

func integerReplacement2(n int) int {
	count := 0
	for n > 1 {
		if n%2 == 0 {
			n /= 2

		} else if (n+1)%4 != 0 || n == 3 {
			n = n - 1

		} else {
			n = n + 1

		}

		count++
	}

	return count

}
