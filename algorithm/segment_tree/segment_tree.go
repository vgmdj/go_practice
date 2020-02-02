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

// Build build the segment tree
func (st *SegmentTree) Build() {
	st.build(0, len(st.data)-1, 0)
}

func (st *SegmentTree) build(l, r, rt int) {
	if l == r {
		st.sum[rt] = st.data[l]
		return
	}

	mid := (l + r) >> 1
	st.build(l, mid, rt<<1|1)
	st.build(mid+1, r, (rt+1)<<1)

	st.buildRoot(rt)

}

// buildRoot update the node info
// if the root is rt, the left child is rt*2+1 = rt<<1+1 = rt<<1|1
// and the right child is rt*2+2 = (rt+1)*2 = (rt+1)<<1
func (st *SegmentTree) buildRoot(rt int) {
	st.sum[rt] = st.sum[rt<<1|1] + st.sum[(rt+1)<<1]
}

// Update add the num at the index i
func (st *SegmentTree) Update(i, num int) {
	st.update(0, len(st.data)-1, i, num, 0)
}

// Update st.data[l] += num
// then update the sum from left to right
func (st *SegmentTree) update(l, r, i, num, rt int) {
	if l == r {
		st.sum[rt] += num
		return
	}

	mid := (l + r) >> 1
	if i <= mid {
		st.update(l, mid, i, num, rt<<1|1)

	} else {
		st.update(mid+1, r, i, num, (rt+1)<<1)

	}

	st.buildRoot(rt)

}

// SumRange return the sum of data from left to right
func (st *SegmentTree) SumRange(l, r int) int {
	return st.sumRange(0, len(st.data)-1, l, r, 0)
}

// L and R represent the data range
// l and r represent the want range
func (st *SegmentTree) sumRange(L, R, l, r, rt int) int {
	// no common interval
	if L > r || R < l {
		return 0
	}

	// the data range all in want range range
	if L >= l && R <= r {
		return st.sum[rt]
	}

	mid := (L + R) >> 1

	return st.sumRange(L, mid, l, r, rt<<1|1) +
		st.sumRange(mid+1, R, l, r, (rt+1)<<1)

}
