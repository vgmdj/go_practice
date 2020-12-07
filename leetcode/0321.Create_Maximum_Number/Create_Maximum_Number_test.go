package Create_Maximum_Number

import (
	"reflect"
	"testing"
)

func Test_maxNumber(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
		k     int
	}
	tests := []struct {
		name    string
		args    args
		wantRes []int
	}{
		// TODO: Add test cases.
		{
			"case",
			args{
				[]int{6, 7, 5},
				[]int{4, 8, 1},
				3,
			},
			[]int{8, 7, 5},
		},
		{
			"case",
			args{
				[]int{6, 7},
				[]int{6, 0, 4},
				5,
			},
			[]int{6, 7, 6, 0, 4},
		},
		{
			"case",
			args{
				[]int{3, 5, 6, 4, 9},
				[]int{9, 2, 8, 3, 6, 4, 5, 8},
				4,
			},
			[]int{9, 9, 8, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := maxNumber(tt.args.nums1, tt.args.nums2, tt.args.k); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("maxNumber() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
