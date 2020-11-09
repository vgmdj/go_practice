package Sort_Integers_by_The_Number_of_1_Bits

/*
我们定义 bit[i] 为二进制表示下数字 1 的个数，则可以列出递推式：
bit[i] = bit[i>>1] + (i \& 1)
bit[i]=bit[i>>1]+(i&1)

*/

var bitNumber = [1e+4 + 1]int{}

func init() {
	for i := 0; i < len(bitNumber); i++ {
		bitNumber[i] = bitNumber[i>>1] + (i & 1)
	}
}

func sortByBits(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	left, right := 0, len(arr)-1
	mid, next := arr[0], 1
	for left < right {
		if largerThan(arr[next], mid) {
			arr[next], arr[right] = arr[right], arr[next]
			right--

		} else {
			arr[next], arr[left] = arr[left], arr[next]
			left++
			next++

		}
	}

	sortByBits(arr[:next-1])
	sortByBits(arr[next:])

	return arr
}

func largerThan(a, b int) bool {
	c1, c2 := bitCount(a), bitCount(b)
	if c1 > c2 {
		return true
	}

	if c1 < c2 {
		return false
	}

	return a > b

}

func bitCount(n int) int {
	return bitNumber[n]

}
