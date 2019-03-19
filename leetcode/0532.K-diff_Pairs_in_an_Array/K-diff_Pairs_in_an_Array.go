package K_diff_Pairs_in_an_Array

import "sort"

//use map , time O(n), space O(n)
func findPairs(nums []int, k int) int {
	if k < 0 {
		return 0
	}

	var count int
	var m = make(map[int]int, len(nums))
	for _, value := range nums {
		m[value]++
	}

	for key, _ := range m {
		if k == 0 && m[key] > 1 {
			count++
			continue
		}

		if k > 0 && m[key+k] > 0 {
			count++
		}
	}

	return count

}

//use tow pointers time O(nlogn), space O(1)
func findPairs2(nums []int, k int) int {
	if len(nums) < 2 {
		return 0
	}
	sort.Ints(nums)

	p, q, result, lastq := 0, 1, 0, nums[0]-1

	for q < len(nums) {
		if p == q {
			q++
			continue
		}

		if nums[q]-nums[p] < k {
			q++
			continue

		}

		if nums[q]-nums[p] > k {
			p++
			continue

		}

		if nums[q] != lastq {
			result++
			lastq = nums[q]
		}

		p++
		q++

	}

	return result

}
