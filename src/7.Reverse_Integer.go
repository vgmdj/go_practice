package main

import (
	"log"
)

// min = -1 << 31 + 1
var MIN int = -0x80000000

// max = 1 << 31 -1
var MAX int = 0x7FFFFFFF

func reverse(x int) int {

	result := 0

	for x != 0 {
		remainder := x % 10
		x = x / 10
		result = result*10 + remainder

	}

	if result < MIN || result > MAX {
		result = 0
	}

	return result

}

func main() {
	log.Println(reverse(-1230))
}
