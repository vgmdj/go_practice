package Bitwise_AND_of_Numbers_Range


/*
临位的与, 从某位后, 就会互为相反数, 相反数与的结果为0
根据这一规律可以进行简化

2-5
2 0010
3 0011
4 0100
5 0101

结果为, 每位的最长相同前缀,后面补0
如上 2-5, 最小前缀为0
若为 4-5, 最小前缀为01, 后被2个0

m 最小, 最小前缀前的0最多, 即 m有最小前缀, 其余值都是更大的前缀
n 最大, 最小前缀前1的位置最远, 即 n有包含最小前缀的最长部分,

取m与n 前缀相同部分

*/
func rangeBitwiseAnd(m int, n int) int {
	var offset = uint(0)

	for m != 0 && n != m {
		n >>= 1
		m >>= 1
		offset++
	}

	return int(m << offset)

}
