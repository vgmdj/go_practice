package Reverse_String_II

func reverseStr(s string, k int) string {
	runes := []rune(s)

	for i := 0; i < len(runes)-1; i += 2 * k {
		right := i + k - 1
		if i+k > len(runes)-1 {
			reverseRune(runes, i, len(runes)-1)
			break
		}
		reverseRune(runes, i, right)
	}

	return string(runes)

}

func reverseRune(str []rune, left, right int) {
	for ; left < right; left, right = left+1, right-1 {
		str[left], str[right] = str[right], str[left]
	}
}
