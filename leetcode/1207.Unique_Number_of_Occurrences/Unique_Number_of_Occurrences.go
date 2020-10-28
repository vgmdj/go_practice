package Unique_Number_of_Occurrences

func uniqueOccurrences(arr []int) bool {
	counter := make(map[int]int)
	for _, v := range arr {
		counter[v]++
	}

	flag := make(map[int]bool)
	for _, v := range counter {
		if flag[v] {
			return false
		}
		flag[v] = true
	}

	return true
}
