package Happy_Number

func isHappy1(n int) bool {
	var numbers = make(map[int]bool)

	for n != 1 && !numbers[n] {
		numbers[n] = true

		var sum int
		for ; n != 0; n /= 10 {
			tmp := n % 10
			sum += tmp * tmp
		}

		n = sum
	}

	return n == 1

}

func isHappy(n int) bool {
	if n == 1 || n == 7 {
		return true
	}
	if n < 10 {
		return false
	}
	temp := n
	p := 0
	for temp > 0 {
		d := temp % 10
		p = p + d*d
		temp = temp / 10
	}
	return isHappy(p)

}
