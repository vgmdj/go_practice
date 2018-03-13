package main

import "log"

func removeElement(nums []int, val int) int {
	var head, tail int
	for tail = len(nums) - 1; head <= tail; {
		if nums[head] == val {
			nums[head], nums[tail] = nums[tail], nums[head]
			tail--
		} else {
			head++
		}

	}

	return head
}

func main() {
	nums := []int{3, 3, 3, 3, 4, 3}
	length := removeElement(nums, 3)
	log.Println(nums, length)

}
