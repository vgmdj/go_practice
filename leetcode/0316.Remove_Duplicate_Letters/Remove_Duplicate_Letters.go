package Remove_Duplicate_Letters

func removeDuplicateLetters(s string) string {
	stack, cnt := make([]byte, 0), [26]int{}
	for _, v := range s {
		cnt[v-'a']++
	}

	inStack := [26]bool{}
	for i := range s {
		ch := s[i]

		for len(stack) > 0 && !inStack[ch-'a'] && ch < stack[len(stack)-1] && cnt[stack[len(stack)-1]-'a'] > 0 {
			inStack[stack[len(stack)-1]-'a'] = false
			stack = stack[:len(stack)-1]

		}

		if !inStack[ch-'a'] {
			stack = append(stack, ch)
			inStack[ch-'a'] = true
		}

		cnt[ch-'a']--
	}

	return string(stack)

}
