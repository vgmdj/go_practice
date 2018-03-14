package main

/*
使用的是异或操作，满足交换率和结合率，
相同异或为0
0与任何数相异或为其本身

*/
func singleNumber(nums []int) int {
	result := 0

	for _, v := range nums {
		result = result ^ v

	}

	return result

}
