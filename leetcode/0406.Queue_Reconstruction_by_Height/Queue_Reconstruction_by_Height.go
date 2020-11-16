package Queue_Reconstruction_by_Height

import "sort"

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		a, b := people[i], people[j]
		return a[0] > b[0] || a[0] == b[0] && a[1] < b[1]
	})

	result := make([][]int, len(people))
	for _, person := range people {
		idx := person[1]
		copy(result[idx+1:], result[idx:])
		result[idx] = person
	}

	return result
}
