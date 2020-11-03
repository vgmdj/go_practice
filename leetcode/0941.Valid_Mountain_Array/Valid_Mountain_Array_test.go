package Valid_Mountain_Array

import "testing"

func Test_validMountainArray(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			"case",
			args{
				[]int{1, 5, 3},
			},
			true,
		},
		{
			"case",
			args{
				[]int{3, 5, 3},
			},
			true,
		},
		{
			"case",
			args{
				[]int{1, 5, 5},
			},
			false,
		},
		{
			"case",
			args{
				[]int{1, 5, 6, 7, 8, 10, 9, 4, 3, 2, 1},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validMountainArray(tt.args.A); got != tt.want {
				t.Errorf("validMountainArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
