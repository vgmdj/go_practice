package Reverse_Words_in_a_String

import (
	"strings"
)

func reverseWords1(s string) string {
	words := strings.Fields(s)

	left, right := 0, len(words)-1
	for left < right {
		words[left], words[right] = words[right], words[left]
		left++
		right--
	}

	return strings.Join(words, " ")
}
