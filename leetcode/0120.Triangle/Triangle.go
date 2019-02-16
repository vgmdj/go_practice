package _120_Triangle

import (
	"math"
)

func minimumTotal(triangle [][]int) int {
	if len(triangle) < 0 {
		return 0
	}

	if len(triangle) == 1 {
		return triangle[0][0]
	}

	min := math.MaxInt32
	for row := 1; row < len(triangle); row++ {
		for column := 0; column < len(triangle[row]); column++ {
			if column == 0 {
				triangle[row][column] += triangle[row-1][column]

			} else if column == len(triangle[row])-1 {
				triangle[row][column] += triangle[row-1][column-1]

			} else {
				triangle[row][column] = Min(triangle[row][column]+triangle[row-1][column-1],
					triangle[row][column]+triangle[row-1][column])

			}

			if row == len(triangle)-1 {
				min = Min(min, triangle[row][column])
			}

		}

	}

	return min

}

func Min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
