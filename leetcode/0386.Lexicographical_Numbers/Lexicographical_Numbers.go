package Lexicographical_Numbers

func lexicalOrder(n int) []int {
	var res []int
	for i := 1; i < 10; i++ {
		helper(i, n, &res)
	}

	return res
}

func helper(cur, n int, res *[]int) {
	if cur > n {
		return
	}

	*res = append(*res, cur)

	for i := 0; i < 10; i++ {
		if cur*10+i > n {
			break
		}

		helper(cur*10+i, n, res)
	}

}
