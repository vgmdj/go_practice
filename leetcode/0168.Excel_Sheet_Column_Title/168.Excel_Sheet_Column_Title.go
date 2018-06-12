package Excel_Sheet_Column_Title

func convertToTitle(n int) string {
	res := ""
	for n > 0 {
		res = string((n-1)%26+65) + res
		n = (n - 1) / 26
	}
	return res
}

func convertToTitle2(n int) string {
	var (
		runes []rune
	)

	for n != 0 {
		current := n % 26
		if current == 0 {
			n--
			runes = append(runes, 'Z')
		} else {
			runes = append(runes, rune(current-1)+'A')
		}

		n = n / 26

	}

	return revert(runes)

}

func revert(runes []rune) string {
	if len(runes) <= 1 {
		return string(runes)
	}

	left, right := 0, len(runes)-1

	for left < right {
		runes[left], runes[right] = runes[right], runes[left]
		right--
		left++
	}

	return string(runes)

}
