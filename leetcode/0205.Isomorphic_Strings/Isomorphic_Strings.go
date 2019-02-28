package Isomorphic_Strings

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var ref = make(map[byte]byte)
	var visited = make([]bool, 128)
	for k := range s {
		pre, ok := ref[s[k]]
		if !ok {
			if visited[t[k]] {
				return false
			}

			ref[s[k]] = t[k]
			visited[t[k]] = true
			continue
		}

		if pre != t[k] {
			return false
		}

	}

	return true

}
