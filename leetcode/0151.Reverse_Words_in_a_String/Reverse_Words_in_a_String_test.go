package Reverse_Words_in_a_String

import "testing"

func Test_reverseWords(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			"space",
			args{
				"  ",
			},
			"",
		},
		{
			"single",
			args{
				"theskyisblue",
			},
			"theskyisblue",
		},
		{
			"normal",
			args{
				"the sky is blue",
			},
			"blue is sky the",
		},
		{
			"trim",
			args{
				"  hello world!  ",
			},
			"world! hello",
		},
		{
			"multiple space",
			args{
				"a good   example",
			},
			"example good a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseWords(tt.args.s); got != tt.want {
				t.Errorf("reverseWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
