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
		if str[i] < '0' || str[i] > '9' {
			break

		}

		if critical, ok := isOverflow(rst*sign, int(str[i]-'0')); ok {
			return critical
		}

		rst = rst*10 + int(str[i]-'0')
	}

	return sign * rst

}

func isOverflow(rst int, add int) (int, bool) {
	if rst > math.MaxInt32/10 || (rst == math.MaxInt32/10 && add > math.MaxInt32%10) {
		return math.MaxInt32, true

	}

	if rst < math.MinInt32/10 || (rst == math.MinInt32/10 && -add < math.MinInt32%10) {
		return math.MinInt32, true
	}

	return 0, false
}
