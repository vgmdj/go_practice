package Number_of_Boomerangs

func numberOfBoomerangs(points [][]int) int {

	result := 0

	for i := 0; i < len(points); i++ {
		boomerang := make(map[int]int, 0)

		for j := 0; j < len(points); j++ {
			if i == j {
				continue
			}

			d := dist(points[i], points[j])
			if _, ok := boomerang[d]; ok {
				result += 2
			}

			boomerang[d]++

		}

	}

	return result

}

func dist(a, b []int) int {
	return (a[0]-b[0])*(a[0]-b[0]) + (a[1]-b[1])*(a[1]-b[1])

}
