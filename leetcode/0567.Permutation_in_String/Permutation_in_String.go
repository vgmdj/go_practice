package Permutation_in_String

func checkInclusion(s1 string, s2 string) bool {
	if len(s2) < len(s1) {
		return false
	}

	s1map := make([]int, 26)
	s2map := make([]int, 26)
	count := 0

	for i := range s1 {
		s1map[s1[i]-'a']++
		s2map[s2[i]-'a']++
	}

	for i := range s1map {
		if s1map[i] == s2map[i] {
			count++
		}
	}

	for i := 0; i < len(s2)-len(s1); i++ {
		if count == 26 {
			return true
		}

		r := s2[i+len(s1)] - 'a'
		l := s2[i] - 'a'

		s2map[r]++
		if s2map[r] == s1map[r] {
			count++

		} else if s2map[r] == s1map[r]+1 {
			count--
		}

		s2map[l]--
		if s2map[l] == s1map[l] {
			count++

		} else if s2map[l] == s1map[l]-1 {
			count--

		}

	}

	return count == 26

}

func checkInclusion1(s1 string, s2 string) bool {
	visited := make([]int, 26)

	for _, v := range s1 {
		visited[v-'a']++
	}

	l, r := 0, 0
	for i := 0; i < len(s2) && l < len(s2) && r < len(s2); i++ {
		//访问值扣除
		visited[s2[i]-'a']--
		r++

		//窗口右边界右移
		if visited[s2[i]-'a'] >= 0 {
			if r-l == len(s1) {
				return true
			}

			continue

		}

		//如果访问值不足，则从左边界到右边界访问值依次补回
		//如果中间有和当前值相同的，则将左边界移到这个位置
		j := l
		for ; j < r; j++ {
			visited[s2[j]-'a']++
			if i != j && s2[j] == s2[i] {
				l = j + 1
				break
			}

		}

		if j+1 != l {
			l = r
		}

	}

	return false

}
