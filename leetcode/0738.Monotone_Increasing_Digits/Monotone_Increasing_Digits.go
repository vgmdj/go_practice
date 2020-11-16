package Monotone_Increasing_Digits

import "math"

func monotoneIncreasingDigits(N int) int {
	result, n := 0, N
	list := make([]int, 10)
	i := 9
	for ; i > 0 && n != 0; i-- {
		list[i] = n % 10
		n /= 10
	}

	start, end := i, 9
	for ; i < len(list)-1; i++ {
		if list[i] > list[i+1] {
			list[i]--
			end = i
			break
		}

	}

	for index := end - 1; index >= start; index-- {
		if list[index] > list[index+1] {
			list[index]--
			continue
		}

		end = index + 1
		break
	}

	for ; start <= end; start++ {
		result = result*10 + list[start]
	}

	rest := int(math.Pow10(9 - end))
	result *= rest
	result += rest - 1

	return result
}
