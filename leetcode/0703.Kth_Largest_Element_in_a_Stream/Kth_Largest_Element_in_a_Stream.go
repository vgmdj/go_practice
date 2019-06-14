package Kth_Largest_Element_in_a_Stream

type KthLargest struct {
	heap []int
	k    int
}

// construct a k min-heap
// 总体的思路就是在k个数里用堆找出最小的值
// 具体讲就是：
// 先对给定的数组构建堆，然后剔除到剩下k个数
// 构建一个大小为k的最小堆，每次新加入的值，和根节点对比
// 如果和根节点相等或比它小，则第k大的值，还是根节点的值
// 如果比根节点的值大，则将其赋给根节点，再调整
func Constructor(k int, nums []int) KthLargest {

	for i := len(nums)/2 - 1; i >= 0; i-- {
		minHeapify(nums, i, len(nums)-1)
	}

	for i := len(nums) - 1; i > 0 && i > k-1; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		minHeapify(nums, 0, i-1)
	}

	if len(nums) >= k {
		nums = nums[:k]
	}

	return KthLargest{
		heap: nums,
		k:    k,
	}

}

func (this *KthLargest) Add(val int) int {

	if len(this.heap) < this.k {
		this.heap = append(this.heap, val)
		minHeapify(this.heap, 0, len(this.heap)-1)
		return this.heap[0]
	}

	if val <= this.heap[0] {
		return this.heap[0]
	}

	this.heap[0] = val
	minHeapify(this.heap, 0, len(this.heap)-1)

	return this.heap[0]

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
