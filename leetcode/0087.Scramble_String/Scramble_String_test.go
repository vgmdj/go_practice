package Scramble_String

import "testing"

func Test_isScramble(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		//{
		//	"case",
		//	args{
		//		s1: "abcde",
		//		s2: "caebd",
		//	},
		//	false,
		//},
		{
			"case",
			args{
				s1: "great",
				s2: "rgate",
			},
			true,
		},
		{
			"case",
			args{
				s1: "abcde",
				s2: "cadeb",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isScramble(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("isScramble() = %v, want %v", got, tt.want)
			}
		})
	}
}
