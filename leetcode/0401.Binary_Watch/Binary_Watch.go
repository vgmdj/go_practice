package Binary_Watch

import "fmt"

func readBinaryWatch(num int) []string {
	result := make([]string, 0)
	backTracking(num, &result, 0, 0, 0)
	return result
}

func backTracking(num int, result *[]string, top int, bottom int, pos int) {
	if num < 0 || num > 10 || pos > 10 || top > 11 || bottom > 59 {
		return
	}

	if num == 0 {
		*result = append(*result, fmt.Sprintf("%d:%02d", top, bottom))
		return
	}

	for i := pos; i < 10; i++ {
		if i < 4 {
			backTracking(num-1, result, top|1<<uint(i), bottom, i+1)

		} else {
			backTracking(num-1, result, top, bottom|1<<uint(i-4), i+1)

		}

	}

}

func readBinaryWatch2(num int) []string {

	if num < 0 {
		return []string{}
	}

	res := make([]string, 0)
	for h := 0; h < 12; h++ {
		for m := 0; m < 60; m++ {
			if countingBit(h)+countingBit(m) == num {
				s := fmt.Sprintf("%d:%02d", h, m)
				res = append(res, s)
			}
		}
	}

	return res
}

func countingBit(num int) int {
	res := 0
	for index := 0; index < 6; index++ {
		if num&1 == 1 {
			res++
		}
		num >>= 1
	}
	return res
}
