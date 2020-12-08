package Split_Array_into_Consecutive_Subsequences

import "testing"

func Test_isPossible(t *testing.T) {
	type args struct {
		nums []int
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
				[]int{1, 2, 3, 4, 4, 5},
			},
			false,
		},
		{
			"case",
			args{
				[]int{1, 2, 3, 3, 4, 4, 5, 5},
			},
			true,
		},
		{
			"case",
			args{
				[]int{1, 2, 4},
			},
			false,
		},
		{
			"case",
			args{
				[]int{1, 2, 3, 3, 4, 5},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPossible(tt.args.nums); got != tt.want {
				t.Errorf("isPossible() = %v, want %v", got, tt.want)
			}
		})
	}
}
