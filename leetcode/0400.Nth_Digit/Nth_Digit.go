package Nth_Digit

import "strconv"

func findNthDigit(n int) int {
	charSize, baseNum, period := 1, 9, 1

	for n > charSize*baseNum*period {
		n -= charSize * baseNum * period
		charSize++
		period *= 10

	}

	return int(strconv.Itoa(period + (n-1)/charSize)[(n-1)%charSize] - '0')

}
