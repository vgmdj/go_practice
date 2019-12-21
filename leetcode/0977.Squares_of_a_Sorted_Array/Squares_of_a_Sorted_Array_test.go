package Squares_of_a_Sorted_Array

import (
	"reflect"
	"testing"
)

func Test_sortedSquares(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			"case1",
			args{
				[]int{-4, -3, -2, -1, 0, 1, 1, 2, 2, 3, 5, 9, 10},
			},
			[]int{0, 1, 1, 1, 4, 4, 4, 9, 9, 16, 25, 81, 100},
		},
		{
			"case2",
			args{
				[]int{-7, -3, 2, 3, 11},
			},
			[]int{4, 9, 9, 49, 121},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedSquares(tt.args.A); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedSquares() = %v, want %v", got, tt.want)
			}
		})
	}
}
