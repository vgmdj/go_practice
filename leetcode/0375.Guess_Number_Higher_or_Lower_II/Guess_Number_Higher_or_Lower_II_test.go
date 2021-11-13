package Guess_Number_Higher_or_Lower_II

import "testing"

func Test_getMoneyAmount(t *testing.T) {
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
				10,
			},
			16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMoneyAmount(tt.args.n); got != tt.want {
				t.Errorf("getMoneyAmount() = %v, want %v", got, tt.want)
			}
		})
	}
}
