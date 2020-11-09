package Sort_Integers_by_The_Number_of_1_Bits

import (
	"reflect"
	"testing"
)

func Test_sortByBits(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			"case",
			args{
				arr: []int{9, 1, 2, 4, 10, 15, 2, 3, 5, 6},
			},
			[]int{1, 2, 2, 4, 3, 5, 6, 9, 10, 15},
		},
		{
			"case",
			args{
				[]int{0, 1, 2, 3, 4, 5, 6, 7, 8},
			},
			[]int{0, 1, 2, 4, 8, 3, 5, 6, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortByBits(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortByBits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bitCount(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			"case",
			args{
				1,
			},
			1,
		},
		{
			"case",
			args{
				3,
			},
			2,
		},
		{
			"case",
			args{
				8,
			},
			1,
		},
		{
			"case",
			args{
				16,
			},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bitCount(tt.args.n); got != tt.want {
				t.Errorf("bitCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
