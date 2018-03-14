package main

import "log"

func twoSum(nums []int, target int) []int {
	sum := make(map[int]int)
	result := []int{}
	for k, v := range nums {
		mapk := target - v
		mapv, ok := sum[v]
		if !ok {
			sum[mapk] = k
			continue
		}

		result = append(result, k)
		result = append(result, mapv)

		return result

	}

	return nil

}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	log.Println(twoSum(nums, target))
}
