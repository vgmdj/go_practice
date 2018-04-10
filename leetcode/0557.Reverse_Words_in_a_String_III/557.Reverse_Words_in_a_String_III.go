package Reverse_Words_in_a_String_III

import (
	"strings"
)

func reverseWords(s string) string {
	runes := []rune(s)
	var start int

	for i := 0; i < len(runes); i++ {
		if runes[i] == ' ' && i > 0 {
			reverseRunes(runes, start, i-1)
			start = i + 1
		}

		if i == len(runes)-1 {
			reverseRunes(runes, start, i)
			start = i + 1
		}

	}

	return string(runes)

}

func reverseRunes(runes []rune, left, right int) {
	for ; left < right; left, right = left+1, right-1 {
		runes[left], runes[right] = runes[right], runes[left]
	}
}

//another method
func reverseWordsII(s string) string {
	ss := strings.Split(s, " ")
	for i, s := range ss {
		ss[i] = revers(s)
	}
	return strings.Join(ss, " ")
}

func revers(s string) string {
	bytes := []byte(s)
	i, j := 0, len(bytes)-1
	for i < j {
		bytes[i], bytes[j] = bytes[j], bytes[i]
		i++
		j--
	}
	return string(bytes)
}
