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
		sum:  make([]int, len(array)<<2),
	}
}

// PushUp update the node info
// if the root is rt, the left child is rt*2 = rt<<1
// and the right child is rt*2+1 = rt<<1+1 = rt<<1|1
func (st *SegmentTree) PushUp(rt int) {
	st.sum[rt] = st.sum[rt<<2] + st.sum[rt<<2|1]
}

// Build build the segment tree
func (st *SegmentTree) Build(l, r, rt int) {
	if l == r {
		st.sum[rt] = st.data[l]
		return
	}

	mid := (l + r) >> 1
	st.Build(l, mid, rt<<1)
	st.Build(mid+1, r, rt<<1|1)

	st.PushUp(rt)

}

// Update st.data[l] += num
// then update the sum from left to right
func (st *SegmentTree) Update(num, l, r, rt int) {

}

// Sum return the sum of data from left to right
func (st *SegmentTree) Sum(l, r int) int {

	return 0
}
