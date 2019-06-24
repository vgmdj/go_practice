package Reshape_the_Matrix

func matrixReshape(nums [][]int, r int, c int) [][]int {
	if len(nums) == 0 {
		return nums
	}

	or, oc := len(nums), len(nums[0])
	if or*oc != c*r {
		return nums
	}

	result := make([][]int, 0)
	for i := 0; i < r; i++ {
		result = append(result, make([]int, c))
	}

	for i := 0; i < c*r; i++ {
		result[i/c][i%c] = nums[i/oc][i%oc]
	}

	return result

}
