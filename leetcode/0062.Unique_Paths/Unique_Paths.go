package Unique_Paths

//0 00 ; 1 01; 2 02
//3 10 ; 4 12; 5 13
//6 20 ; 7 21; 8 22
//
//0 00; 1 01
//2 10; 3 11
//4 20; 5 21
func uniquePaths(m int, n int) int {
	if m <= 0 || n <= 0 {
		return 0
	}

	var paths = make([]int, m*n)
	for row := 0; row < m; row++ {
		for column := 0; column < n; column++ {
			if row == 0 || column == 0 {
				paths[row*n+column] = 1
			} else {
				paths[row*n+column] = paths[(row-1)*n+column] + paths[row*n+(column-1)]
			}

		}
	}

	return paths[m*n-1]

}

//递归写法，超时
func uniquePaths1(m int, n int) int {
	if m == 0 || n == 0 {
		return 0
	}

	if m == 1 || n == 1 {
		return 1
	}

	return uniquePaths1(m-1, n) + uniquePaths1(m, n-1)

}
