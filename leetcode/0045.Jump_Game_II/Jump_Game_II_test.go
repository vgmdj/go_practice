package Jump_Game_II

import "testing"

func Test_jump(t *testing.T) {
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
			"case",
			args{
				[]int{1, 2},
			},
			1,
		},
		{
			"case1",
			args{
				[]int{7, 0, 9, 6, 9, 6, 1, 7, 9, 0, 1, 2, 9, 0, 3},
			},
			2,
		},
		{
			"case2",
			args{
				[]int{2, 3, 1},
			},
			1,
		},
		{
			"case3",
			args{
				[]int{2, 3, 1, 1, 4},
			},
			2,
		},
		{
			"case4",
			args{
				[]int{1, 2, 3},
			},
			2,
		},
		{
			"case5",
			args{
				[]int{1, 1, 1, 1},
			},
			3,
		},
		{
			"case6",
			args{
				[]int{1, 1, 1, 1, 1, 1},
			},
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := jump(tt.args.nums); got != tt.want {
				t.Errorf("jump() = %v, want %v", got, tt.want)
			}
		})
	}
}
