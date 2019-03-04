package Reverse_Vowels_of_a_String

func reverseVowels(s string) string {
	p, q := 0, len(s)-1
	var bts = []byte(s)
	for p < q {
		for ; p < q; p++ {
			if isVowel(bts[p]) {
				break
			}
		}

		for ; p < q; q-- {
			if isVowel(bts[q]) {
				break
			}
		}

		bts[p], bts[q] = bts[q], bts[p]
		p++
		q--
	}

	return string(bts)

}

func isVowel(b byte) bool {
	vowels := []byte{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}
	for _, v := range vowels {
		if b == v {
			return true
		}
	}

	return false

}
