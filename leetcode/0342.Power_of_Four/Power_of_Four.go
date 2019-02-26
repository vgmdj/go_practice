package Power_of_Four

func isPowerOfFour(num int) bool {
	return num > 0 && (num-1)&num == 0 && (num&0x55555555) != 0
}
