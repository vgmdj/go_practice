package How_Many_Numbers_Are_Smaller_Than_the_Current_Number

func smallerNumbersThanCurrent(nums []int) []int {
	result := make([]int, len(nums))
	cnt := [101]int{}
	for _, v := range nums {
		cnt[v]++
	}

	for i := 1; i < len(cnt)-1; i++ {
		cnt[i] += cnt[i-1]
	}

	for i, v := range nums {
		if v > 0 {
			result[i] = cnt[v-1]
		}

	}

	return result
}
