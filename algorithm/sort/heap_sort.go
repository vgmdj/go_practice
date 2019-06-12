package sort

import (
	"container/heap"
)

/*
堆排序分为两个部分，第一个部分是初始化，第二个部分是堆恢复

初始化：
	从最后一个非叶子结点开始进行调整，每次调整后，如果以当前结点为根结点，则这个结点也是以这个结点为根的堆的根结点

调整：
	将根结点与最后一个结点进行互换，同时，将除最后一个结点外的部分进行堆调整恢复

之所以和最后一个结点进行进行互换，是因为，这样和初始化的最后一步类似，可以复用过程

*/

func HeapSort(values []int) {
	heapInit(values)
	//for i := len(values) - 1; i > 0; i-- {
	//	values[0], values[i] = values[i], values[0]
	//	minHeapify(values, 0, i-1)
	//}

	for i := len(values) - 1; i > 0; i-- {
		values[0], values[i] = values[i], values[0]
		maxHeapify(values, 0, i-1)
	}

}

func heapInit(values []int) {
	for i := len(values)/2 - 1; i >= 0; i-- {
		//minHeapify(values, i, len(values)-1)
		maxHeapify(values, i, len(values)-1)
	}
}

func maxHeapify(nums []int, start, end int) {
	father := start
	child := 2*father + 1

	for child <= end {
		if child+1 <= end && nums[child] < nums[child+1] {
			child++
		}

		if nums[father] >= nums[child] {
			return
		}

		nums[father], nums[child] = nums[child], nums[father]
		father = child
		child = 2*father + 1

	}

}

func minHeapify(nums []int, start, end int) {
	root := start
	node := 2*root + 1

	for node <= end {
		if node+1 <= end && nums[node+1] < nums[node] {
			node++
		}

		if nums[root] <= nums[node] {
			return
		}

		nums[root], nums[node] = nums[node], nums[root]
		root = node
		node = 2*root + 1

	}

}

type heapInt []int

func (h heapInt) Len() int {
	return len(h)
}

func (h heapInt) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h heapInt) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h *heapInt) Pop() interface{} {
	tail := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return tail

}

func (h *heapInt) Push(x interface{}) {
	*h = append(*h, x.(int))

}

func HeapSort2(values []int) {
	hp := heapInt{}
	for i := range values {
		heap.Push(&hp, values[i])
	}

	for i := range values {
		values[i] = heap.Pop(&hp).(int)
	}

}
