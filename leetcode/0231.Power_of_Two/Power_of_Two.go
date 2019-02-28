package Power_of_Two

func isPowerOfTwo(n int) bool {
	return n > 0 && ((n & (n - 1)) == 0)

}

func isPowerOfTwo2(n int) bool {
	power := 1
	sth := 0

	for power <= n && power > 0 {
		if power == n {
			return true
		}

		power = power << 1
		sth++
	}

	return false

}
