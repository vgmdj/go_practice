package Smallest_String_With_A_Given_Numeric_Value

func getSmallestString(n int, k int) string {
	result := make([]byte, n)
	index := n - 1

	// 设置 z
	for ; k >= (26 + n - 1); index, n, k = index-1, n-1, k-26 {
		result[index] = 'z'
	}

	if index < 0 {
		return string(result)
	}

	// z 前一个字母
	mod := byte((k - n) % 26)
	result[index] = mod + 'a'
	index--
	n--
	k -= int(mod + 1)

	for ; index >= 0; index-- {
		result[index] = 'a'
	}

	return string(result)
}
