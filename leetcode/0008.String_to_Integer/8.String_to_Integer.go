package String_to_Integer

import (
	"math"
	"strings"
)

func myAtoi(str string) int {
	str = strings.Trim(str, " ")
	if str == "" {
		return 0
	}

	sign, index := 1, 0
	if str[0] == '+' {
		index++
	} else if str[0] == '-' {
		sign = -1
		index++
	}

	rst := 0
	for i := index; i < len(str); i++ {
		if str[i] < '0' || str[i] > '9' || rst > math.MaxInt32 {
			break

		}
		rst = rst*10 + int(str[i]-'0')
	}

	result, over := isOverflow(sign, rst)
	if over {
		return result
	}

	return sign * rst

}

func isOverflow(sign, rst int) (int, bool) {
	if sign == 1 && rst > math.MaxInt32 {
		return math.MaxInt32, true

	}

	if sign == -1 && sign*rst < math.MinInt32 {
		return math.MinInt32, true
	}

	return 0, false
}
