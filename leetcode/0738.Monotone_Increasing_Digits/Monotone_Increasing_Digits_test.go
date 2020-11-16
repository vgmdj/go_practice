package Monotone_Increasing_Digits

import "testing"

func Test_monotoneIncreasingDigits(t *testing.T) {
	type args struct {
		N int
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
				114191537,
			},
			113999999,
		},
		{
			"case",
			args{
				4321,
			},
			3999,
		},
		{
			"case",
			args{
				13332,
			},
			12999,
		},
		{
			"case",
			args{
				4523,
			},
			4499,
		},
		{
			"case",
			args{
				10,
			},
			9,
		},
		{
			"case",
			args{
				4586,
			},
			4579,
		},
		{
			"case",
			args{
				1234,
			},
			1234,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := monotoneIncreasingDigits(tt.args.N); got != tt.want {
				t.Errorf("monotoneIncreasingDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
