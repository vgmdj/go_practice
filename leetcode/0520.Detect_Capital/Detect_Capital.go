package Detect_Capital

func detectCapitalUse(word string) bool {
	if len(word) < 2 {
		return true
	}

	if isLower(word[0]) && isUpper(word[1]) {
		return false
	}

	for i := 2; i < len(word); i++ {
		if isLower(word[1]) && !isLower(word[i]) {
			return false
		}

		if isUpper(word[1]) && !isUpper(word[i]) {
			return false
		}
	}

	return true

}

func isLower(b byte) bool {
	if b >= 'a' && b <= 'z' {
		return true
	}

	return false
}

func isUpper(b byte) bool {
	if b >= 'A' && b <= 'Z' {
		return true
	}

	return false
}
