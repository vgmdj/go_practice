package Group_Anagrams

func groupAnagrams(strs []string) [][]string {
	var result [][]string
	var ref = make(map[string]int)

	for k, v := range strs {
		str := sortStr(v)
		index, ok := ref[str]
		if !ok {
			result = append(result, []string{strs[k]})
			ref[str] = len(result) - 1
			continue
		}

		result[index] = append(result[index], strs[k])
	}

	return result

}

func sortStr(str string) string {
	var bts = []byte(str)
	quickSort(bts)

	return string(bts)

}

func quickSort(data []byte) {
	if len(data) < 2 {
		return
	}

	var left, right = 0, len(data) - 1
	var mid, flag = 1, data[0]

	for left < right {
		if data[mid] <= flag {
			data[left], data[mid] = data[mid], data[left]
			left++
			mid++

		} else {
			data[right], data[mid] = data[mid], data[right]
			right--

		}

	}

	quickSort(data[:mid-1])
	quickSort(data[mid:])

}
