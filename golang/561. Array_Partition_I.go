package main

import (
	"sort"
	"log"
)

func arrayPairSum(nums []int) int {
	sort.Ints(nums)

	log.Println(nums)

	var sum int
	for k,v := range nums{
		if k % 2 == 1{
			log.Println(v)
			sum += v
		}
	}

	return sum

}

func main(){
	var nums = []int{1,4,3,2}
	log.Println(arrayPairSum(nums))

}
