package main

import (
	"log"
)

//the most simple way
func maxSubArray1(nums []int) int {
	max := 0
	if len(nums) > 0 {
		max = nums[0]
	}

	for i := 0; i < len(nums); i++ {
		result := nums[i]

		if result > max {
			max = result
		}

		for j := i + 1; j < len(nums); j++ {
			result += nums[j]
			if result > max {
				max = result
			}

		}

	}

	return max
}

var MIN_INT32 = -0x80000000
var MAX_INT32 = 0xffffffff

//better way to solve the problem
func maxSubArray(nums []int) int {
	max := MIN_INT32
	sum := 0

	for _, v := range nums {
		sum += v

		if max < sum {
			max = sum
		}

		if sum < 0 {
			sum = 0
		}

	}
	return max
}

func main() {
	//nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	nums := []int{-2}
	log.Println(maxSubArray(nums))

}
