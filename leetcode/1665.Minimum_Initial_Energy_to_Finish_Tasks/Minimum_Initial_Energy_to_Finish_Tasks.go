package Minimum_Initial_Energy_to_Finish_Tasks

import "sort"

func minimumEffort(tasks [][]int) int {
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i][1]-tasks[i][0] > tasks[j][1]-tasks[j][0]
	})

	res, curMin := 0, 0
	for i := range tasks {
		if curMin < tasks[i][1] {
			curMin = tasks[i][1]
		}
		res += tasks[i][0]
		curMin -= tasks[i][0]
	}

	res += curMin

	return res
}
