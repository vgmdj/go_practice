package Longest_Valid_Parentheses

import "testing"

func Test_longestValidParentheses(t *testing.T) {
	type args struct {
		s string
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
				"(()",
			},
			2,
		},
		{
			"case2",
			args{
				"(()()(",
			},
			4,
		},
		{
			"case3",
			args{
				"(((((",
			},
			0,
		},
		{
			"case4",
			args{
				"))()(())(()))",
			},
			10,
		},
		{
			"case5",
			args{
				"(((((()()))))",
			},
			12,
		},
		{
			"case6",
			args{
				"()(()",
			},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestValidParentheses(tt.args.s); got != tt.want {
				t.Errorf("longestValidParentheses() = %v, want %v", got, tt.want)
			}
		})
	}
}
