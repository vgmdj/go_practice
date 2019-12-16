package Maximum_Swap

import "strconv"

//string deal
func maximumSwap2(num int) int {
	str := []byte(strconv.Itoa(num))
	start := 0
	position := make(map[byte]int)
	for i := range str {
		position[str[i]] = i
	}

	max := byte('9')
	for ; start < len(str); start++ {
		if str[start] >= max {
			continue
		}

		for max >= '0' {
			if max == str[start] {
				break
			}

			p, ok := position[max]
			if ok && p == start {
				break
			}

			if ok && p > start {
				str[start], str[p] = str[p], str[start]
				result, _ := strconv.Atoi(string(str))
				return result
			}

			max--

			// !ok

		}

	}

	result, _ := strconv.Atoi(string(str))
	return result
}
