package Gray_Code

/*
递归生成码表
这种方法基于格雷码是反射码的事实，利用递归的如下规则来构造：

- 1位格雷码有两个码字
- (n+1)位格雷码中的前2n个码字等于n位格雷码的码字，按顺序书写，加前缀0
- (n+1)位格雷码中的后2n个码字等于n位格雷码的码字，按逆序书写，加前缀1 [4]
- n+1位格雷码的集合 = n位格雷码集合(顺序)加前缀0 + n位格雷码集合(逆序)加前缀1

2位格雷码	3位格雷码	4位格雷码	4位自然二进制码
00			000			0000		0000
01			001			0001		0001
11			011			0011		0010
10			010			0010		0011
			110			0110		0100
			111			0111		0101
			101			0101		0110
			100			0100		0111
						1100		1000
						1101		1001
						1111		1010
						1110		1011
						1010		1100
						1011		1101
						1001		1110
						1000		1111

ps:若要提升效率，可以将 0 ，1 ，2， 3， 4， 5 ... 位的情况直接返回
情况列出越多，效率越高

	if n <= 0 {
		return []int{0}
	}

	if n == 1 {
		return []int{0, 1}
	}

    if n == 2 {
		return []int{0, 1, 3, 2}
	}

	if n == 3 {
		return []int{0, 1, 3, 2, 6, 7, 5, 4}
	}

	if n == 4 {
		return []int{0, 1, 3, 2, 6, 7, 5, 4, 12, 13, 15, 14, 10, 11, 9, 8}
	}

	......

*/

func grayCode(n int) []int {
	if n <= 0 {
		return []int{0}
	}

	gcp := grayCode(n - 1)

	result := make([]int, len(gcp)*2)

	copy(result, gcp)

	base := 1 << uint(n-1)
	for i := len(gcp) - 1; i >= 0; i-- {
		result[len(result)-i-1] = gcp[i] | base
	}

	return result

}

func grayCode2(n int) []int {
	return helper(n, 1, []int{0})
}

func helper(n int, base int, result []int) []int {
	if n == 0 {
		return result
	}

	length := len(result)
	for i := length - 1; i >= 0; i-- {
		result = append(result, result[i]|base)
	}

	return helper(n-1, base*2, result)
}
