package main

import (
	"log"
)

//bitsum%3 要么是0,要么是1,因为数组结构组成是3n+1
func singleNumberII(nums []int) int {
	var rst int = 0
	var i uint8

	for i = 0; i < 32; i++ {
		bitSum := 0
		for _, v := range nums {
			if (v & (1 << i)) > 0 {
				bitSum++
			}

		}

		rst = rst | ((bitSum % 3) << i)
	}

	return int(int32(rst))
}

func main() {
	nums := []int{-2, -2, 1, 1, -3, 1, -3, -3, -4, -2}
	log.Println(singleNumberII(nums))

}
