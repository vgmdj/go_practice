package Increasing_Decreasing_String

import "testing"

func Test_sortString(t *testing.T) {
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
			"case",
			args{
				"aaaabbbbcccc",
			},
			"abccbaabccba",
		},
		{
			"case",
			args{
				"rat",
			},
			"art",
		},
		{
			"case",
			args{
				"leetcode",
			},
			"cdelotee",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortString(tt.args.s); got != tt.want {
				t.Errorf("sortString() = %v, want %v", got, tt.want)
			}
		})
	}
}
