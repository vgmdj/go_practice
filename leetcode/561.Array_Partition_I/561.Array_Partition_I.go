package Array_Partition_I

import (
	"sort"
)

func arrayPairSum(nums []int) int {
	sort.Ints(nums)

	var sum int
	for k, v := range nums {
		if k%2 == 1 {
			sum += v
		}
	}

	return sum

}
