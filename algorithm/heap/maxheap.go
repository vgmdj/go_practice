package heap

// MaxHeap the max heap
type MaxHeap struct {
}

// NewMaxHeap return the object of max heap
func NewMaxHeap() *MaxHeap {
	return &MaxHeap{}
}

// Heapify heapify the data
func (mh *MaxHeap) Heapify(data []int, start, end int) {
	father := start
	child := 2*father + 1

	for child <= end {
		if child+1 <= end && data[child+1] > data[child] {
			child = child + 1
		}

		if data[child] <= data[father] {
			return
		}

		data[father], data[child] = data[child], data[father]
		father = child
		child = 2*father + 1

	}

}
