package assigncookies

import "sort"

func findContentChildren(g []int, s []int) int {
	result := 0
	sort.Ints(g)
	sort.Ints(s)

	index, c := 0, 0
	for c < len(g) && index < len(s) {
		if s[index] >= g[c] {
			result++
			c++
		}

		index++

	}

	return result
}
