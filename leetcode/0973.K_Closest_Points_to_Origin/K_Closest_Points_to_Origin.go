package K_Closest_Points_to_Origin

type distHelper struct {
	dist []int
}

func newDistHelper(points [][]int) *distHelper {
	dist := make([]int, len(points))
	for i := 0; i < len(points); i++ {
		dist[i] = distSquare(points[i])
	}

	return &distHelper{
		dist: dist,
	}
}

func (dh *distHelper) heapInit(points [][]int) {
	for i := len(points)/2 - 1; i >= 0; i-- {
		dh.minHeapify(points, i, len(points)-1)
	}
}

func (dh *distHelper) minHeapify(points [][]int, start, end int) {
	father := start
	child := father*2 + 1
	for child <= end {
		if child+1 <= end && dh.dist[child] > dh.dist[child+1] {
			child = child + 1
		}

		if dh.dist[father] <= dh.dist[child] {
			return
		}

		points[father], points[child] = points[child], points[father]
		dh.swap(father, child)
		father = child
		child = father*2 + 1

	}

}

func (dh *distHelper) swap(p1, p2 int) {
	dh.dist[p1], dh.dist[p2] = dh.dist[p2], dh.dist[p1]
}

func kClosest(points [][]int, K int) [][]int {
	if len(points) <= K {
		return points
	}
	dh := newDistHelper(points)
	dh.heapInit(points)
	for i := len(points) - 1; i >= len(points)-K; i-- {
		points[0], points[i] = points[i], points[0]
		dh.swap(0, i)
		dh.minHeapify(points, 0, i-1)
	}

	return points[len(points)-K:]
}

func distSquare(point []int) int {
	if len(point) < 2 {
		return 0
	}

	return point[0]*point[0] + point[1]*point[1]
}
