package How_Many_Numbers_Are_Smaller_Than_the_Current_Number

import (
	"reflect"
	"testing"
)

func Test_smallerNumbersThanCurrent(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			"100",
			args{
				nums: []int{1, 50, 100},
			},
			[]int{0, 1, 2},
		},
		{
			"7777",
			args{
				nums: []int{7, 7, 7, 7},
			},
			[]int{0, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallerNumbersThanCurrent(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("smallerNumbersThanCurrent() = %v, want %v", got, tt.want)
			}
		})
	}
}
