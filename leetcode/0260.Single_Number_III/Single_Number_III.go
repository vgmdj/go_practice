package Single_Number_III

/*
所有数异或后, 结果就等于两个只出现一次的数的异或
结果里的每位1, 都表示 这两个数在这位上一个为0, 一个为1

取结果里的这个1, 其余的数, 因都出现过两次, 在归类时都会被归到同一种情况里
所以, 最终可以得出这两个数


*/
func singleNumber(nums []int) []int {
	a, b, xor, mask := 0, 0, 0, 0x1

	for _, v := range nums {
		xor ^= v
	}

	for xor&0x1 == 0 {
		xor >>= 1
		mask <<= 1
	}

	for _, v := range nums {
		if v&mask == 0 {
			a ^= v

		} else {
			b ^= v

		}
	}

	return []int{a, b}

}
