package Compare_Version_Numbers

import "strconv"

func compareVersion(version1 string, version2 string) int {
	index1, index2 := 0, 0
	for index1 < len(version1) || index2 < len(version2) {
		var v1, v2 int
		v1, index1 = getNextVersion(version1, index1)
		v2, index2 = getNextVersion(version2, index2)

		if v1 > v2 {
			return 1
		}

		if v1 < v2 {
			return -1
		}
		index1++
		index2++

	}

	return 0

}

func getNextVersion(version string, start int) (int, int) {
	if start >= len(version) {
		return 0, start + 1
	}

	trim := func(str string) int {
		left, right := 0, len(str)-1
		for left < right && str[left] == '0' {
			left++
		}

		result, _ := strconv.Atoi(str[left : right+1])

		return result
	}

	end := start
	for ; end < len(version); end++ {
		if version[end] == '.' {
			break
		}
	}

	return trim(version[start:end]), end

}
