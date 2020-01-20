package segment_tree

// SegmentTree the data structure define
type SegmentTree struct {
	data []int
	sum  []int
}

// NewSegmentTree the construct of the segment tree
func NewSegmentTree(array []int) *SegmentTree {
	return &SegmentTree{
		data: array,
		sum:  make([]int, len(array)*4),
	}
}

// PushUp update the node info
func (st *SegmentTree) PushUp(rt int) {

}

// Build build the segment tree
func (st *SegmentTree) Build(l, r, rt int) {

}

// Update st.data[l] += num
// then update the sum from left to right
func (st *SegmentTree) Update(num, l, r, rt int) {

}

// Sum return the sum of data from left to right
func (st *SegmentTree) Sum(l, r int) int {

	return 0
}
