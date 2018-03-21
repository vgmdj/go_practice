package Reverse_String

func reverseString1(s string) string {
	if len(s) <= 1 {
		return s
	}

	bytes := []byte(s)
	left := 0
	right := len(s) - 1
	for left < right {
		bytes[left], bytes[right] = bytes[right], bytes[left]
		left++
		right--

	}

	return string(bytes)

}

func reverseString(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
