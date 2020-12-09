package Split_Array_into_Fibonacci_Sequence

import (
	"reflect"
	"testing"
)

func Test_splitIntoFibonacci(t *testing.T) {
	type args struct {
		S string
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
				"1101111",
			},
			[]int{11, 0, 11, 11},
		},
		{
			"case",
			args{
				"11235813",
			},
			[]int{1, 1, 2, 3, 5, 8, 13},
		},
		{
			"case",
			args{
				"0123",
			},
			[]int{},
		},
		{
			"case",
			args{
				"0112",
			},
			[]int{0, 1, 1, 2},
		},
		{
			"case",
			args{
				"112358130",
			},
			[]int{},
		},
		{
			"case",
			args{
				"123456579",
			},
			[]int{123, 456, 579},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitIntoFibonacci(tt.args.S); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitIntoFibonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}
