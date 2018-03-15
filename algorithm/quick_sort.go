package main

import (
	"log"
)

func quickSort(values []int) {
	if len(values) <= 1 {
		return
	}
	mid, i := values[0], 1
	head, tail := 0, len(values)-1
	for head < tail {
		log.Println(values)
		if values[i] > mid {
			values[i], values[tail] = values[tail], values[i]
			tail--
		} else {
			values[i], values[head] = values[head], values[i]
			head++
			i++
		}
	}
	values[head] = mid
	quickSort(values[:head])
	quickSort(values[head+1:])

}

func main() {
	nums := []int{7, 6, 5, 4, 3, 2, 1}

	quickSort(nums)

	log.Println(nums)

}
