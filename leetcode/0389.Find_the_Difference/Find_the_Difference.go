package Find_the_Difference

/*
原理有点像之前的找单独数字
相同的数异或会为0,s与t中只多出一个字符,可用此方式求解

方法二
使用visited数组计数, [26]int{}
先遍历t,对visited++, 如遇到a 则 visited['a'-'a']++
再遍历s,对visited--, 如遇到a 则 visited['a'-'a']--

最后只会有一个位置上的数不为0, 如0
结果就是 0+'a'
*/
func findTheDifference(s string, t string) byte {
	var result byte

	for _, v := range s {
		result ^= byte(v)
	}

	for _, v := range t {
		result ^= byte(v)
	}

	return result

}

func findTheDifference2(s string, t string) byte {
	var visited = make([]int, 26)
	for _, v := range t {
		visited[byte(v)-'a']++
	}

	for _, v := range s {
		visited[byte(v)-'a']--
	}

	for k, v := range visited {
		if v != 0 {
			return byte(k + 'a')
		}
	}

	return 0

}
