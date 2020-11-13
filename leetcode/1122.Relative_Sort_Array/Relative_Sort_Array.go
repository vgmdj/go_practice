package Relative_Sort_Array

func relativeSortArray(arr1 []int, arr2 []int) []int {
	upper := 0
	for _, v := range arr1 {
		if v > upper {
			upper = v
		}
	}
	frequency := make([]int, upper+1)
	for _, v := range arr1 {
		frequency[v]++
	}

	ans := make([]int, 0, len(arr1))
	for _, v := range arr2 {
		for ; frequency[v] > 0; frequency[v]-- {
			ans = append(ans, v)
		}
	}
	for v, freq := range frequency {
		for ; freq > 0; freq-- {
			ans = append(ans, v)
		}
	}
	return ans
}
