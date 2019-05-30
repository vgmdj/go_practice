package Sliding_Window_Maximum

import "container/list"

type Node struct {
	Index int
	Value int
}

func maxSlidingWindow2(nums []int, k int) []int {
	if len(nums) == 0 || k < 1 || k > len(nums) {
		return nil
	}

	l := list.New()

	result := make([]int, len(nums))
	for i := range nums {
		if l.Front() != nil && i-k >= l.Front().Value.(Node).Index {
			l.Remove(l.Front())
		}

		for l.Len() != 0 && l.Back().Value.(Node).Value < nums[i] {
			l.Remove(l.Back())
		}

		l.PushBack(Node{
			Index: i,
			Value: nums[i],
		})

		result[i] = l.Front().Value.(Node).Value

	}

	return result[k-1:]

}
