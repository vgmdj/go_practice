package Add_Digits

func addDigits(num int) int {
	if num == 0 {
		return 0
	}

	return (num-1)%9 + 1

}

func addDigits2(num int) int {

	for num >= 10 {
		sum := 0

		for num >= 10 {
			sum += num % 10
			num /= 10
		}

		num += sum

	}

	return num

}
