package Wildcard_Matching

// time limit
func isMatch2(s string, p string) bool {
	return backTracking(s, p, 0, 0)
}

func backTracking(s, p string, pi, si int) bool {
	if len(s) <= si && len(p) <= pi {
		return true
	}

	if len(p) <= pi {
		return false
	}

	switch p[pi] {
	case '*':
		if pi == len(p)-1 {
			return true
		}

		for i := 0; i <= len(s)-si; i++ {
			if backTracking(s, p, pi+1, si+i) {
				return true
			}
		}

	default:
		if si >= len(s) || (p[pi] != '?' && p[pi] != s[si]) {
			return false
		}

		return backTracking(s, p, pi+1, si+1)

	}

	return false
}
