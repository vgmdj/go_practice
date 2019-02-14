package _3Sum_Closest

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		return sum(nums)
	}

	sort.Ints(nums)

	closest := math.MaxInt32
	for first := 0; first <= len(nums)-3; first++ {
		second, third := first+1, len(nums)-1
		for second < third {
			sum := nums[first] + nums[second] + nums[third]
			if sum == target {
				return target
			} else if sum < target {
				second++
			} else {
				third--
			}

			closest = closer(target, closest, sum)
		}
	}

	return closest

}

func sum(nums []int) int {
	var result int

	for _, v := range nums {
		result += v
	}

	return result
}

func closer(target, pre, next int) int {
	result1 := target - pre
	result2 := target - next

	if result1 < 0 {
		result1 = -result1
	}

	if result2 < 0 {
		result2 = -result2
	}

	if result1 < result2 {
		return pre
	}

	return next

}
