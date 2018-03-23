package Roman_to_Integer

var roman = make(map[byte]int)

func init() {
	roman = map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

}

func romanToInt(s string) int {
	res := 0

	last := 0
	for i := len(s) - 1; i >= 0; i-- {
		temp := roman[s[i]]

		sign := 1
		if temp < last {
			//小数在大数的左边，要减去小数
			sign = -1
		}

		res += sign * temp

		last = temp
	}

	return res
}

//another way
func romanToIntII(s string) int {
	if len(s) == 0 {
		return -1
	}

	var rst int
	var nextI byte

	for i := len(s) - 1; i >= 0; i-- {
		switch s[i] {
		default:
			return -1

		case 'I':
			if nextI == 'V' || nextI == 'X' {
				rst--
			} else {
				rst++
			}

		case 'X':
			if nextI == 'L' || nextI == 'C' {
				rst -= 10
			} else {
				rst += 10
			}

		case 'C':
			if nextI == 'D' || nextI == 'M' {
				rst -= 100
			} else {
				rst += 100
			}

		case 'V':
			rst += 5

		case 'L':
			rst += 50

		case 'D':
			rst += 500

		case 'M':
			rst += 1000

		}

		nextI = s[i]

	}

	return rst

}

//the best way to solved this problem
func romanToIntIII(s string) int {
	sum := 0
	lastC := '0'
	for i, c := range s {
		if i != 0 {
			lastC = rune(s[i-1])
		}
		if c == 'I' {
			sum = sum + 1
		}
		if c == 'V' {
			sum = sum + 5 - checkPre(lastC, 'I')
		}
		if c == 'X' {
			sum = sum + 10 - checkPre(lastC, 'I')
		}
		if c == 'L' {
			sum = sum + 50 - checkPre(lastC, 'X')
		}
		if c == 'C' {
			sum = sum + 100 - checkPre(lastC, 'X')
		}
		if c == 'D' {
			sum = sum + 500 - checkPre(lastC, 'C')
		}
		if c == 'M' {
			sum = sum + 1000 - checkPre(lastC, 'C')
		}
	}
	return sum
}

func checkPre(r1, r2 rune) int {
	if r1 != r2 {
		return 0
	}
	if r1 == 'I' {
		return 1 << 1

	} else if r1 == 'X' {
		return 10 << 1

	} else if r1 == 'C' {
		return 100 << 1

	}
	return 0
}
