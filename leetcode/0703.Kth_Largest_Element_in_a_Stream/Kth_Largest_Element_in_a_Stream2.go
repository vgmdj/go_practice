package Kth_Largest_Element_in_a_Stream

import "container/heap"

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	last := (*h)[:len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return last
}

// real thing starts here
type KthLargest2 struct {
	Nums *IntHeap
	K    int
}

func Constructor2(k int, nums []int) KthLargest2 {
	h := &IntHeap{}
	heap.Init(h)
	// push all elements to the heap
	for i := 0; i < len(nums); i++ {
		heap.Push(h, nums[i])
	}
	// remove the smaller elements from the heap such that only the k largest elements are in the heap
	for len(*h) > k {
		heap.Pop(h)
	}
	return KthLargest2{h, k}
}

func (this *KthLargest2) Add(val int) int {
	if len(*this.Nums) < this.K {
		heap.Push(this.Nums, val)
	} else if val > (*this.Nums)[0] {
		heap.Push(this.Nums, val)
		heap.Pop(this.Nums)
	}
	return (*this.Nums)[0]
}
