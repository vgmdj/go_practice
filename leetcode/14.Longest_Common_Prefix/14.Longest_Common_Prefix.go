package Longest_Common_Prefix

func longestCommonPrefix(strs []string) string {
	if len(strs) < 1 {
		return ""
	} else if len(strs) < 2 {
		return strs[0]
	}

	var rst, count = strs[0], 0

	for i := 0; i < len(rst); i++ {
		sameCount := 0
		for j := 1; j < len(strs); j++ {
			if len(strs[j]) > i && strs[j][i] == rst[i] {
				sameCount++
			}
		}

		if sameCount != len(strs)-1 {
			break
		}

		count++
	}

	return rst[:count]

}
