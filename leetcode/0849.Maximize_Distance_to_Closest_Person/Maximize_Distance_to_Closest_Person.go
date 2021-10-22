package Maximize_Distance_to_Closest_Person

func maxDistToClosest(seats []int) int {
	result := 0
	left, right := 0, len(seats)-1
	for ; left <= right; left++ {
		if seats[left] == 1 {
			result = max(result, left)
			break
		}
	}

	for ; right >= left; right-- {
		if seats[right] == 1 {
			result = max(result, len(seats)-right-1)
			break
		}
	}

	for i := left; i <= right && left <= right; i++ {
		if seats[i] == 1 {
			left = i
			continue
		}

		result = max(result, (i-left-1)/2+1)
	}

	return result

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
