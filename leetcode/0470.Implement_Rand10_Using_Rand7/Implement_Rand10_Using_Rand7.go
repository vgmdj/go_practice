package Implement_Rand10_Using_Rand7

func rand10() int {
	sum := 0
	row := 0
	col := 0
	for {
		row = rand7()
		col = rand7()
		sum = (row-1)*7 + col
		if sum <= 40 {
			break
		}
	}
	return 1 + (sum-1)%10
}
