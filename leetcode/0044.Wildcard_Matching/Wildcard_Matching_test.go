package Wildcard_Matching

import "testing"

func Test_isMatch(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			"case1",
			args{
				"aa",
				"a",
			},
			false,
		},
		{
			"case2",
			args{
				"aa",
				"*",
			},
			true,
		},
		{
			"case3",
			args{
				"aaa",
				"a?a",
			},
			true,
		},
		{
			"case4",
			args{
				"cb",
				"?a",
			},
			false,
		},
		{
			"case5",
			args{
				"adceb",
				"*a*b",
			},
			true,
		},
		{
			"case6",
			args{
				"acdcb",
				"a*c?b",
			},
			false,
		},
		{
			"case7",
			args{
				"",
				"**",
			},
			true,
		},
		{
			"case8",
			args{
				"ho",
				"ho**",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatch(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("isMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
