package Rotate_Image

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
2*2
00-10 10-11 11-01

3*3
00-20 20-22 22-02
01-10 10-21 21-12


4*4
00-30 30-33 33-03
01-20 20-32 32-13
02-10 10-31 31-23



*/

func TestRotateImage(t *testing.T) {
	ast := assert.New(t)

	case1 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	rotate(case1)
	ast.Equal([][]int{
		{7, 4, 1},
		{8, 5, 2},
		{9, 6, 3},
	}, case1)

	case2 := [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}
	rotate(case2)
	ast.Equal([][]int{
		{15, 13, 2, 5},
		{14, 3, 4, 1},
		{12, 6, 8, 9},
		{16, 7, 10, 11},
	}, case2)

}
