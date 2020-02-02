package Largest_Rectangle_in_Histogram

func largestRectangleArea1(heights []int) int {
	if len(heights) == 0 {
		return 0
	}

	st := NewSegment(heights)
	st.Build(0, len(heights)-1, 0)

	return divide(st, heights, 0, len(heights)-1)

}

func divide(st *Segment, heights []int, left, right int) int {
	if left > right {
		return -1
	}

	if left == right {
		return heights[left]
	}

	mm := st.RangeMin(left, right)
	ma := heights[mm] * (right - left + 1)

	la := divide(st, heights, left, mm-1)
	ra := divide(st, heights, mm+1, right)

	return Max(ma, la, ra)

}

type Segment struct {
	data []int
	min  []int
}

func NewSegment(nums []int) *Segment {
	return &Segment{
		data: nums,
		min:  make([]int, len(nums)*4),
	}
}

func (st *Segment) Build(left, right, rt int) {
	if left == right {
		st.min[rt] = left
		return
	}

	mid := (left + right) >> 1
	st.Build(left, mid, rt<<1|1)
	st.Build(mid+1, right, (rt+1)<<1)

	st.buildRoot(rt)

}

func (st *Segment) buildRoot(rt int) {
	left, right := st.min[rt<<1|1], st.min[(rt+1)<<1]
	min := left
	if st.data[right] < st.data[left] {
		min = right
	}

	st.min[rt] = min

}

func (st *Segment) RangeMin(left, right int) int {
	return st.rangeMin(left, right, 0, len(st.data)-1, 0)
}

func (st *Segment) rangeMin(left, right, L, R, rt int) int {
	if right < L || left > R {
		return -1
	}

	if left <= L && right >= R {
		return st.min[rt]
	}

	mid := (L + R) >> 1

	lp := st.rangeMin(left, right, L, mid, rt<<1|1)
	rp := st.rangeMin(left, right, mid+1, R, (rt+1)<<1)

	if lp == -1 || rp == -1 {
		return lp + rp + 1
	}

	if st.data[lp] < st.data[rp] {
		return lp
	}

	return rp
}

func Max(nums ...int) int {
	r := nums[0]
	for i := 1; i < len(nums); i++ {
		if r < nums[i] {
			r = nums[i]
		}
	}

	return r

}
