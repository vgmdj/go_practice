package Merge_Intervals

import "sort"

type Interval struct {
	Start int
	End   int
}

func merge(intervals []Interval) []Interval {
	if len(intervals) < 2 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})

	var result = []Interval{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		if intervals[i].Start <= result[len(result)-1].End {
			result[len(result)-1].End = Max(intervals[i].End, result[len(result)-1].End)
			continue
		}

		result = append(result, intervals[i])
	}

	return result

}

func Max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
