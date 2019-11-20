package Substring_with_Concatenation_of_All_Words

func findSubstring(s string, words []string) []int {
	if len(words) == 0 {
		return []int{}
	}

	length := len(words[0])
	size := len(words) * length

	result := make([]int, 0)
	counter := make(map[string]int)
	for _, v := range words {
		counter[v]++
	}

	left := 0
	for i := 0; i <= len(s)-length; {
		selector := s[i : i+length]
		c, ok := counter[selector]
		if ok && c > 0 {
			counter[selector]--
			if i-left+length == size {
				result = append(result, left)
			}
			i += length

			continue
		}

		alreadyIncr := false
		for left <= i-length {
			counter[s[left:left+length]]++
			left += length
			if s[left-length-1:left] == selector {
				alreadyIncr = true
				break
			}
		}

		if !alreadyIncr {
			i++
			left = i
		}

	}

	return result
}
