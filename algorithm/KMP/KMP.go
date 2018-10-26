package kmp

func kmpIndexOf(s string, m string) int {
	if s == "" || m == "" || len(s) < len(m) {
		return -1
	}

	var (
		si, mi int

		next = getNextArray(m)
	)

	for mi < len(m) && si < len(s) {
		if s[si:si+1] == m[mi:mi+1] {
			si++
			mi++
		} else if next[mi] == -1 {
			si++
		} else {
			mi = next[mi]
		}
	}

	if mi == len(m) {
		return si - mi
	}

	return -1

}

func getNextArray(m string) []int {
	if len(m) == 1 {
		return []int{-1}
	}

	var (
		next = make([]int, len(m))

		pos = 2
		cn  = 0
	)

	next[0] = -1
	next[1] = 0

	for pos < len(m) {
		if m[pos-1:pos] == m[cn:cn+1] {
			cn++
			next[pos] = cn
			pos++
		} else if cn > 0 {
			cn = next[cn]
		} else {
			next[pos] = 0
			pos++
		}

	}

	return next
}
