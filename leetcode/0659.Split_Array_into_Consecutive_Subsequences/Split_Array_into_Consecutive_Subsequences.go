package Split_Array_into_Consecutive_Subsequences

func isPossible(nums []int) bool {
	cnt := make(map[int]int, 0)
	for _, c := range nums {
		cnt[c]++
	}

	endCnt := make(map[int]int, 0)
	for _, v := range nums {
		if cnt[v] == 0 {
			continue
		}

		if endCnt[v-1] > 0 {
			cnt[v]--
			endCnt[v-1]--
			endCnt[v]++
			continue
		}

		if cnt[v+1] > 0 && cnt[v+2] > 0 {
			cnt[v]--
			cnt[v+1]--
			cnt[v+2]--
			endCnt[v+2]++
			continue
		}

		return false

	}

	return true

}
