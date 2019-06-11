package Restore_IP_Addresses

func restoreIpAddresses(s string) []string {
	results := make([]string, 0)

	helper(s, 0, 0, &results, "")

	return results

}

func helper(s string, index int, count int, results *[]string, subset string) {
	if index == len(s) && count == 4 {
		*results = append(*results, subset)
		return
	}

	if index >= len(s) || count > 4 || (4-count)*3 < len(s)-index || 4-count > len(s)-index {
		return
	}

	if subset != "" {
		subset += "."
	}

	for i := 0; i < 3; i++ {
		if !isValidNum(s, index, index+i+1) {
			continue
		}

		helper(s, index+i+1, count+1, results, subset+s[index:index+1+i])
	}

}

func isValidNum(s string, start, end int) bool {
	if end > len(s) {
		return false
	}

	if end-start !=1 && s[start] == '0' {
		return false
	}

	if end-start == 3 && s[start:end] > "255" {
		return false
	}

	return true
}
