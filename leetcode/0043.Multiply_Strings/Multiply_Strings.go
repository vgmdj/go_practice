package Multiply_Strings

func multiply(num1 string, num2 string) string {
	if len(num1) == 0 || len(num2) == 0 || num1 == "0" || num2 == "0" {
		return "0"
	}

	result := "0"

	for i := len(num2) - 1; i >= 0; i-- {

		for j := byte(0); j < num2[i]-'0'; j++ {
			result = addTwoString(result, num1)
		}

		num1 += "0"

	}

	return result

}

func addTwoString(str1, str2 string) string {
	if len(str1) < len(str2) {
		str1, str2 = str2, str1
	}

	result := make([]byte, len(str1)+1)
	carry, index := byte(0), len(result)-1

	tail1, tail2 := len(str1)-1, len(str2)-1

	for tail1 >= 0 || tail2 >= 0 {
		sum := byte(0)
		n1, n2 := byte(0), byte(0)
		if tail1 >= 0 {
			n1 = str1[tail1] - '0'
			tail1--
		}

		if tail2 >= 0 {
			n2 = str2[tail2] - '0'
			tail2--
		}

		sum = n1 + n2 + carry
		result[index] = sum%10 + '0'
		carry = sum / 10
		index--

	}

	if carry == 1 {
		result[index] = '1'
		index--
	}

	return string(result[index+1:])
}
