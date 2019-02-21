package Keyboard_Row

import "strings"

func findWords(words []string) []string {
	if len(words) == 0 {
		return words
	}

	var res []string
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 2,
		"d": 1,
		"e": 0,
		"f": 1,
		"g": 1,
		"h": 1,
		"i": 0,
		"j": 1,
		"k": 1,
		"l": 1,
		"m": 2,
		"n": 2,
		"o": 0,
		"p": 0,
		"q": 0,
		"r": 0,
		"s": 1,
		"t": 0,
		"u": 0,
		"v": 2,
		"w": 0,
		"x": 2,
		"y": 0,
		"z": 2,
	}

	for _, wv := range words {
		tv := strings.ToLower(wv)
		idx := -1

		if len(tv) == 0 {
			continue
		}

		idx = m[string(tv[0])]
		flag := true
		for _, v := range tv {
			if m[string(v)] != idx {
				flag = false
				break
			}
		}
		if flag {
			res = append(res, wv)
		}
	}
	return res
}
