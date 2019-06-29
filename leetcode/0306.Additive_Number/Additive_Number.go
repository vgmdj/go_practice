package Additive_Number

func isAdditiveNumber(num string) bool {
	for f := 1; f <= len(num)/2; f++ {
		first := num[:f]
		if first != "0" && first[0] == '0' {
			return false
		}

		tail := 0
		for s := 1; f+s < len(num) && tail <= len(num); s++ {
			second := num[f : f+s]
			length := max(len(first), len(second))
			tail = f + s + length

			if backTracking(num, first, second, f+s, tail) || backTracking(num, first, second, f+s, tail+1) {
				return true
			}

		}

	}

	return false
}

func backTracking(num string, first, second string, ts, te int) bool {
	if ts == len(num) {
		return true
	}

	if len(first) == 0 || len(second) == 0 ||
		(first != "0" && first[0] == '0') ||
		(second != "0" && second[0] == '0') ||
		te > len(num) || (num[ts:te] != "0" &&  num[ts] == '0') ||
		!stringAddCompare(first, second, num[ts:te]) {
		return false
	}

	length := max(len(second), te-ts)
	return backTracking(num, second, num[ts:te], te, te+length) || backTracking(num, second, num[ts:te], te, te+length+1)

}

func stringAddCompare(first, second string, next string) bool {
	carry := byte(0)
	ft, st, nt := len(first)-1, len(second)-1, len(next)-1

	for ft >= 0 && st >= 0 {
		result := first[ft] + second[st] - '0' + carry
		carry = 0
		if result > '9' {
			result = result - 10
			carry = 1

		}

		if result != next[nt] {
			return false
		}

		ft, st, nt = ft-1, st-1, nt-1

	}

	for ft >= 0 {
		result := first[ft] + carry
		carry = 0
		if result > '9' {
			result = result - 10
			carry = 1
		}

		if result != next[nt] {
			return false
		}

		nt--
		ft--
	}

	for st >= 0 {
		result := second[st] + carry
		carry = 0
		if result > '9' {
			result = result - 10
			carry = 1
		}

		if result != next[nt] {
			return false
		}

		nt--
		st--
	}

	if carry == 0 && nt == -1 {
		return true
	}

	if carry == 1 && nt == 0 {
		return next[0] == '1'

	}

	return false

}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
