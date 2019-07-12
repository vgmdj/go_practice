package Counting_Bits

func countBits(num int) []int {
	result := make([]int, num+1)

	power := 1
	for i := 1; power <= num; i++ {
		result[power] = 1
		power = 1 << uint(i)

	}

	index := 1
	for i := 3; i <= num; i++ {
		if result[i] == 1 {
			index = 1
			continue
		}

		result[i] = result[index] + 1
		index++

	}

	return result

}
