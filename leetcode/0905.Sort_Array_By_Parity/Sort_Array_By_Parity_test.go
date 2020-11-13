package Sort_Array_By_Parity

import (
	"reflect"
	"testing"
)

func Test_sortArrayByParity(t *testing.T) {
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
			"case",
			args{
				[]int{1, 4},
			},
			[]int{4, 1},
		},
		{
			"case",
			args{
				[]int{3},
			},
			[]int{3},
		},
		{
			"case",
			args{
				[]int{1, 2, 3, 4, 5, 6, 6},
			},
			[]int{6, 2, 6, 4, 5, 3, 1},
		},
		{
			"case",
			args{
				[]int{1, 3},
			},
			[]int{1, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortArrayByParity(tt.args.A); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortArrayByParity() = %v, want %v", got, tt.want)
			}
		})
	}
}
