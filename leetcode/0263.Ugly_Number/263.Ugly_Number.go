package Ugly_Number

func isUgly(num int) bool {
	d := []int{2, 3, 5}
	for i := 0; i < len(d) && num != 0; i++ {
		for num%d[i] == 0 {
			num = num / d[i]
		}
	}
	return num == 1
}

func isUglyII(num int) bool {
	if num == 0 {
		return false
	}

	for num != 1 {
		if num%2 == 0 {
			num /= 2
			continue
		}

		if num%3 == 0 {
			num /= 3
			continue
		}

		if num%5 == 0 {
			num /= 5
			continue
		}

		return false

	}

	return true
}
