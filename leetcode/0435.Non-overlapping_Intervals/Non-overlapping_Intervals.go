package Non_overlapping_Intervals

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	max, right := 1, intervals[0][1]
	for _, p := range intervals[1:] {
		if p[0] >= right {
			max++
			right = p[1]
		}
	}

	return len(intervals) - max

}
