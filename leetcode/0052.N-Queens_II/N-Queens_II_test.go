package N_Queens_II

import "testing"

func Test_totalNQueens(t *testing.T) {
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
			"case1",
			args{
				1,
			},
			1,
		},
		{
			"case2",
			args{
				4,
			},
			2,
		},
		{
			"case3",
			args{
				5,
			},
			10,
		},
		{
			"case4",
			args{
				8,
			},
			92,
		},
		{
			"case5",
			args{
				10,
			},
			724,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := totalNQueens(tt.args.n); got != tt.want {
				t.Errorf("totalNQueens() = %v, want %v", got, tt.want)
			}
		})
	}
}
