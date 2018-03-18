package main

import "log"

//中心点枚举 O(n^2)的时间复杂度
func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}

	var start, length, longest int
	for i := 0; i < len(s); i++ {
		length = findLongestPalindrome(s, i, i)
		if length > longest {
			longest = length
			start = i - length/2
		}

		length = findLongestPalindrome(s, i, i+1)
		if length > longest {
			longest = length
			start = i - length/2 + 1
		}
	}

	return s[start : start+longest]

}

func findLongestPalindrome(s string, left, right int) int {
	var length int
	for left >= 0 && right < len(s) {
		if s[left] != s[right] {
			break
		}

		length = right - left + 1
		left--
		right++
	}
	return length
}

func main() {
	log.Println(longestPalindrome("cbbd"))
}
