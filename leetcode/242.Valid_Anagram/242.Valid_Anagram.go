package Valid_Anagram

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	ss := []rune(s)
	ts := []rune(t)
	counter := make([]int, 26)
	for i := 0; i < len(s); i++ {
		counter[ss[i]-'a']++
		counter[ts[i]-'a']--
	}

	for _, v := range counter {
		if v != 0 {
			return false
		}
	}

	return true
}
