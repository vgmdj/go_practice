package heap

// IHeap the interface of heap
type IHeap interface {
	Heapify(values []int, start, end int)
}

// Heap the structure of heap
type Heap struct {
	h    IHeap
	list []int
}

// NewHeap return the object of max heap
func NewHeap(h IHeap) *Heap {
	return &Heap{
		list: make([]int, 0),
		h:    h,
	}
}

// Push push the value
func (h *Heap) Push(v int) {
	h.list = append(h.list, v)
	h.swap()
	h.downToUp()
}

// Pop pop the value of heap
func (h *Heap) Pop() (int, bool) {
	if h.Len() == 0 {
		return -1, false
	}

	v := h.list[0]
	h.swap()
	h.list = h.list[:h.Len()-1]
	h.upToDown()

	return v, true
}

// Len return the length of max heap
func (h *Heap) Len() int {
	return len(h.list)
}

func (h *Heap) swap() {
	h.list[0], h.list[h.Len()-1] = h.list[h.Len()-1], h.list[0]
}

func (h *Heap) downToUp() {
	length := len(h.list)
	for i := length/2 - 1; i >= 0; i-- {
		h.h.Heapify(h.list, i, length-1)
	}
}

func (h *Heap) upToDown() {
	h.h.Heapify(h.list, 0, len(h.list)-1)

}
