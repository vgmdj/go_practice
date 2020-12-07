package Reorganize_String

func reorganizeString(S string) string {
	n := len(S)
	if n <= 1 {
		return S
	}

	maxCnt := 0
	cnt := [26]int{}
	for _, c := range S {
		cnt[c-'a']++
		if cnt[c-'a'] > maxCnt {
			maxCnt = cnt[c-'a']
		}
	}

	if maxCnt > (n+1)/2 {
		return ""
	}

	result := make([]byte, n)
	oddIndex, evenIndex, halfLen := 1, 0, n/2
	for i, c := range cnt {
		b := byte('a' + i)
		for c > 0 && c <= halfLen && oddIndex < n {
			result[oddIndex] = b
			c--
			oddIndex += 2
		}

		for c > 0 {
			result[evenIndex] = b
			c--
			evenIndex += 2

		}

	}

	return string(result)

}
