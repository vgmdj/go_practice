package Matrix_Cells_in_Distance_Order

import (
	"container/list"
)

func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
	result := make([][]int, 0)
	buckets := make([][][]int, R+C)
	for row := 0; row < R; row++ {
		for col := 0; col < C; col++ {
			dis := distance(row, col, r0, c0)
			buckets[dis] = append(buckets[dis], []int{row, col})
		}
	}

	for _, bucket := range buckets {
		result = append(result, bucket...)
	}

	return result
}

func distance(x1, y1, x2, y2 int) int {
	return abs(x1-x2) + abs(y1-y2)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func allCellsDistOrder2(R int, C int, r0 int, c0 int) [][]int {
	result := make([][]int, 0)

	visited := make([][]bool, R)
	for i := range visited {
		visited[i] = make([]bool, C)
	}

	queue := list.New()

	pushQueue := func(r, c int) {
		if r < 0 || c < 0 || r >= R || c >= C {
			return
		}

		if visited[r][c] {
			return
		}

		visited[r][c] = true
		queue.PushBack([]int{r, c})
	}

	pushQueue(r0, c0)
	for queue.Len() != 0 {
		b := queue.Front().Value.([]int)
		queue.Remove(queue.Front())

		result = append(result, b)
		pushQueue(b[0]-1, b[1])
		pushQueue(b[0]+1, b[1])
		pushQueue(b[0], b[1]-1)
		pushQueue(b[0], b[1]+1)

	}

	return result
}
