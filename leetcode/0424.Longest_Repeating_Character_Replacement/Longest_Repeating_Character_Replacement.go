package Longest_Repeating_Character_Replacement

//sliding window
func characterReplacement(s string, k int) int {
	cnt := make([]int, 26)
	left, right, max := 0, 0, 0

	for ; right < len(s); right++ {
		cnt[s[right]-'A']++
		if max < cnt[s[right]-'A'] {
			max = cnt[s[right]-'A']
		}

		if right-left >= max+k {
			cnt[s[left]-'A']--
			left++
		}

	}

	return right-left

}
