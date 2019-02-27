package Number_of_1_Bits

func hammingWeight(num uint32) int {
	var result int

	for i := 31; i >= 0; i-- {
		result += int((num >> uint(i)) & 1)
	}

	return result

}
