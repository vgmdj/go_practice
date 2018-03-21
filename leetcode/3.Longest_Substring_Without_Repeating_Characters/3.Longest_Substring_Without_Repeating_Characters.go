package Longest_Substring_Without_Repeating_Characters

import (
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	substring := ""
	maxLength := 0

	for _, v := range s {
		substring = removeDuplicate(substring, byte(v))
		if len(substring) > maxLength {
			maxLength = len(substring)
		}
	}

	return maxLength

}

func removeDuplicate(s string, char byte) string {
	index := strings.IndexByte(s, char)
	if index < 0 {
		return s + string(char)
	}

	if index+1 >= len(s) {
		return s[index:]
	}

	return s[index+1:] + string(char)

}
