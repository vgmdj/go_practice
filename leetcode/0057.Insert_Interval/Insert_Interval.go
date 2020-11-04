package Insert_Interval

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(newInterval) == 0 {
		return intervals
	}

	if len(intervals) == 0 {
		return [][]int{newInterval}
	}

	if newInterval[0] > intervals[len(intervals)-1][1] {
		return append(intervals, newInterval)
	}

	index := 0
	result := make([][]int, 0)
	left, right := 0, len(intervals)-1
	for left <= right {
		mid := (left + right) / 2
		if intervals[mid][0] < newInterval[0] {
			left = mid + 1

		} else {
			right = mid - 1

		}

	}

	index = right + 1
	result = append(result, intervals[:index]...)

	insertInterval(&result, newInterval)

	for index < len(intervals) {
		insertInterval(&result, intervals[index])
		index++
	}

	return result

}

func insertInterval(result *[][]int, interval []int) {
	if len(*result) == 0 {
		*result = append(*result, interval)
		return
	}

	last := &(*result)[len(*result)-1]
	if interval[0] > (*last)[1] {
		*result = append(*result, interval)
		return
	}

	(*last)[0] = Min((*last)[0], interval[0])
	(*last)[1] = Max((*last)[1], interval[1])

}

func Min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
