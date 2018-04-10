package Two_Sum

import "testing"

func Test(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	t.Log(twoSum(nums, target))
}
