package Minimum_Number_of_Arrows_to_Burst_Balloons

import (
	"sort"
)

func findMinArrowShots(points [][]int) int {
	if len(points) < 2 {
		return len(points)
	}

	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	top, result := points[0], 1
	for _, point := range points {
		if point[0] >= top[0] && point[0] <= top[1] {
			top[0] = max(point[0], top[0])
			top[1] = min(point[1], top[1])

		} else {
			top[0] = point[0]
			top[1] = point[1]
			result++

		}

	}

	return result

}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
