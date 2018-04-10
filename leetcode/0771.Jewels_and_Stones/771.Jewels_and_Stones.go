package Jewels_and_Stones

import (
	"strings"
)

func numJewelsInStones(J string, S string) int {
	var counter = make(map[byte]int)
	var rst int

	for _, v := range S {
		counter[byte(v)] = counter[byte(v)] + 1
	}

	for _, v := range J {
		rst += counter[byte(v)]
	}

	return rst

}

func numJewelsInStones2(J string, S string) int {
	sum := 0
	for _, v := range J {
		sum += strings.Count(S, string(v))
	}
	return sum
}

func numJewelsInStones3(J string, S string) int {
	var counter = make([]int, 256)
	var rst int

	for _, v := range S {
		counter[v]++
	}

	for _, v := range J {
		rst += counter[v]
	}

	return rst

}
