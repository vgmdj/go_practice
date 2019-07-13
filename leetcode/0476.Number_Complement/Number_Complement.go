package Number_Complement

func findComplement(num int) int {
	result := 0
	index := 31
	for index > 0 {
		if num>>uint(index)&1 == 1 {
			break
		}
		index--
	}

	for index >= 0 {
		b := num>>uint(index)&1 ^ 1

		result += b << uint(index)

		index--

	}

	return result

}
