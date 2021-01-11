package heap

// MinHeap the min heap
type MinHeap struct {
}

// NewMinHeap return the object of min heap
func NewMinHeap() *MinHeap {
	return &MinHeap{}
}

// Heapify min heapify
func (mh *MinHeap) Heapify(data []int, start, end int) {
	father := start
	child := 2*father + 1

	for child <= end {
		if child+1 <= end && data[child+1] < data[child] {
			child = child + 1
		}

		if data[child] >= data[father] {
			return
		}

		data[father], data[child] = data[child], data[father]
		father = child
		child = 2*father + 1

	}

}
