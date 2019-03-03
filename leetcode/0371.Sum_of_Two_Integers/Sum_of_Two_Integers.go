package Sum_of_Two_Integers

/*
a ^ b 结果为无进位的加
每位相加的结果, 0与1 得1, 1与1 得0, 0与0 得0

a & b 结果为每位的进位情况
只有1与1, 才会进位

(a&b)<<1 , 两个1相遇, 前一位才需要加上进位的1

重复过程, 走到没有进位为止

*/
func getSum(a int, b int) int {
	var sum = a ^ b
	var carry = (a & b) << 1

	if carry == 0 {
		return sum
	}

	return getSum(sum, carry)
}
