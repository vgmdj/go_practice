package main

import (
	"log"
	"strings"
)

func isPalindrome(s string) bool {
	if s == "" {
		return true
	}

	s = strings.ToLower(s)
	left := 0
	right := len(s) - 1

	for left <= right {
		if !isAlpha(s[left]) {
			left++
			continue
		}

		if !isAlpha(s[right]) {
			right--
			continue
		}

		if s[left] != s[right] {
			return false
		}

		left++
		right--

	}

	return true

}

func isAlpha(b byte) bool {
	if 'a' <= b && b <= 'z' || 'A' <= b && b <= 'Z' || '0' <= b && b <= '9' {
		return true
	}
	return false
}

func main() {
	log.Println(isPalindrome("A man, a plan, a canal: Panama"))
	log.Println(isPalindrome("race a car"))
}
