package Is_Subsequence

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}

	if len(s) > len(t) {
		return false
	}

	for p, q := 0, 0; p < len(s) && q < len(t); {
		if s[p] == t[q] {
			p++
			q++
		} else {
			q++
		}

		if p == len(s) {
			return true
		}

	}

	return false

}
