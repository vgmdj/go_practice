package Insert_Interval

import (
	"reflect"
	"testing"
)

func Test_insert(t *testing.T) {
	type args struct {
		intervals   [][]int
		newInterval []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{
			"case1",
			args{
				[][]int{},
				[]int{2, 5},
			},
			[][]int{{2, 5}},
		}, {
			"case2",
			args{
				[][]int{},
				[]int{},
			},
			[][]int{},
		}, {
			"case3",
			args{
				[][]int{{3, 5}},
				[]int{},
			},
			[][]int{{3, 5}},
		},
		{
			"case4",
			args{
				[][]int{{1, 2}, {3, 5}, {4, 7}, {6, 8}},
				[]int{2, 5},
			},
			[][]int{{1, 8}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insert(tt.args.intervals, tt.args.newInterval); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insert() = %v, want %v", got, tt.want)
			}
		})
	}
}
