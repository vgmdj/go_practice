package segment_tree

type SegmentTree struct {
	tree  []int
	data  []int
	merge func(int, int) int
}

func NewSegmentTree(array []int, merge func(int, int) int) *SegmentTree {
	return &SegmentTree{
		tree:  make([]int, len(array)*4),
		data:  make([]int, len(array)),
		merge: merge,
	}
}

// build segment_tree from left to right
func (st *SegmentTree) BuildSegmentTree(index, left, right int) int {
	if left == right {
		st.tree[index] = st.data[left]
		return st.tree[index]
	}

	leftChild := 2*index + 1
	rightChild := 2*index + 2
	mid := (left + right) / 2

	leftMid := st.BuildSegmentTree(leftChild, left, mid)
	rightMid := st.BuildSegmentTree(rightChild, mid+1, right)

	st.tree[index] = st.merge(leftMid, rightMid)

	return st.tree[index]

}
