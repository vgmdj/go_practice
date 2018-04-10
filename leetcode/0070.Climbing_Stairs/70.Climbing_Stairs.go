package Climbing_Stairs

//斐波那契数列
func climbStairs(n int) int {
	var (
		first  = 0
		second = 1
		rst    = 0
	)

	for i := 0; i < n; i++ {
		rst = first + second
		first = second
		second = rst

	}

	return rst

}
