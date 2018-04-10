package Plus_One

func plusOne(digits []int) []int {
	if len(digits) == 0 {
		return []int{}
	}

	res := make([]int, len(digits)+1)
	carry := 1
	for lastIndex := len(digits) - 1; lastIndex >= 0; lastIndex-- {
		res[lastIndex+1], carry = plusCarry(digits[lastIndex], carry)

	}

	if carry == 0 {
		return res[1:]
	}

	res[0] = carry
	return res

}

func plusCarry(num, lastC int) (result, carry int) {
	if num+lastC == 10 {
		return 0, 1
	}

	return num + lastC, 0
}

//another way
func plusOneII(digits []int) []int {
	if len(digits) == 0 {
		return []int{}
	}

	digits[len(digits)-1]++
	for i := len(digits) - 1; i > 0; i-- {
		if digits[i] == 10 {
			digits[i] = 0
			digits[i-1]++
		}
	}

	if digits[0] == 10 {
		digits[0] = 0
		return append([]int{1}, digits[:]...)
	}

	return digits

}
