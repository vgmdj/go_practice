package Count_Numbers_with_Unique_Digits

/*
n=1: res=10

n=2: res=9*9+10=91 # 两位数第一位只能为1-9，第二位只能为非第一位的数，加上一位数的所有结果

n=3: res=9 * 9 * 8+91=739 # 三位数第一位只能为1-9，第二位只能为非第一位的数，第三位只能为非前两位的数，加上两位数的所有结果

n=4: res=9 * 9 * 8 * 7+739=5275 # 同上推法

*/

func countNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 1
	}

	res, mul := 10, 9

	for i := 2; i <= n; i++ {
		mul *= 10 - i + 1
		res += mul

	}

	return res

}
