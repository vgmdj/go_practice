package Multiply_Strings

import "testing"

func Test_multiply(t *testing.T) {
	type args struct {
		num1 string
		num2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			"case1",
			args{
				"11",
				"23",
			},
			"253",
		},
		{
			"case2",
			args{
				"2",
				"3",
			},
			"6",
		},
		{
			"case3",
			args{
				"24",
				"11",
			},
			"264",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := multiply(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}
