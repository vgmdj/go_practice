package Word_Pattern

import "strings"

func wordPattern(pattern string, str string) bool {
	var strarr = strings.Split(str, " ")
	if len(pattern) != len(strarr) {
		return false
	}

	var ref = make(map[byte]string)
	var reverse = make(map[string]byte)

	for i := range pattern {
		pre, ok := ref[pattern[i]]
		if !ok {
			ref[pattern[i]] = strarr[i]

		} else if pre != strarr[i] {
			return false

		}

		next, ok := reverse[strarr[i]]
		if !ok {
			reverse[strarr[i]] = pattern[i]

		} else if next != pattern[i] {
			return false

		}

	}

	return true

}
