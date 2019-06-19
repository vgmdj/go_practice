package Add_Strings

func addStrings(num1 string, num2 string) string {
	result := make([]byte, len(num1)+len(num2))
	tail1, tail2, tail := len(num1)-1, len(num2)-1, len(result)-1
	carry := byte(0)

	for ; tail1 >= 0 && tail2 >= 0; tail1, tail2, tail = tail1-1, tail2-1, tail-1 {
		num := num1[tail1] + num2[tail2] + carry - '0'
		carry = 0
		if num > '9' {
			num -= 10
			carry = 1
		}

		result[tail] = num
	}

	for ; tail1 >= 0; tail1, tail = tail1-1, tail-1 {
		num := num1[tail1] + carry
		carry = 0
		if num > '9' {
			num -= 10
			carry = 1
		}

		result[tail] = num
	}

	for ; tail2 >= 0; tail2, tail = tail2-1, tail-1 {
		num := num2[tail2] + carry
		carry = 0
		if num > '9' {
			num -= 10
			carry = 1
		}

		result[tail] = num
	}

	if carry == 1 {
		result[tail] = '1'
		tail--
	}

	return string(result[tail+1:])

}
