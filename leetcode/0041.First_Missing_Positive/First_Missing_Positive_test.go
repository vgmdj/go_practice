package First_Missing_Positive

import "testing"

func Test_firstMissingPositive(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			"same",
			args{
				[]int{1, 1},
			},
			2,
		},
		{
			"simple",
			args{
				[]int{-1,4,2,1,9,10},
			},
			3,
		},
		{
			"negative",
			args{
				[]int{3, 4, -1, 1},
			},
			2,
		},
		{
			"positive",
			args{
				[]int{7, 8, 9, 11, 12},
			},
			1,
		},
		{
			"null",
			args{
				[]int{-3},
			},
			1,
		},
		{
			"zero",
			args{
				[]int{-1, 10000},
			},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstMissingPositive(tt.args.nums); got != tt.want {
				t.Errorf("firstMissingPositive() = %v, want %v", got, tt.want)
			}
		})
	}
}
