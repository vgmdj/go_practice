package Reverse_Words_in_a_String

func reverseWords(s string) string {
	if len(s) == 0 {
		return ""
	}

	bts := []byte(s)
	start := 0

	res := make([]byte, 0, len(bts))
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			continue
		}

		if i != 0 && s[i-1] != ' ' && i != start {
			reverse(bts[start:i])
			res = append(res, bts[start:i]...)
			res = append(res, ' ')
		}

		start = i + 1

	}

	if bts[len(bts)-1] != ' ' {
		if start == 0 {
			return s
		}

		reverse(bts[start:])
		res = append(res, bts[start:]...)
	}

	if len(res) > 0 && res[len(res)-1] == ' ' {
		res = res[:len(res)-1]
	}

	reverse(res)

	return string(res)

}

func reverse(data []byte) {
	left, right := 0, len(data)-1
	for left < right {
		data[left], data[right] = data[right], data[left]
		left++
		right--
	}

	return

}
