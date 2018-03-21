package Single_Number_II

//bitsum%3 要么是0,要么是1,因为数组结构组成是3n+1
func singleNumberII(nums []int) int {
	var rst int = 0
	var i uint8

	for i = 0; i < 32; i++ {
		bitSum := 0
		for _, v := range nums {
			if (v & (1 << i)) > 0 {
				bitSum++
			}

		}

		rst = rst | ((bitSum % 3) << i)
	}

	return int(int32(rst))
}
