package main

func hammingDistance(x int, y int) int {
	var count int
	for x != 0 || y != 0 {
		if x%2 != y%2 {
			count++
		}

		x = x >> 1
		y = y >> 1

	}

	return count

}
