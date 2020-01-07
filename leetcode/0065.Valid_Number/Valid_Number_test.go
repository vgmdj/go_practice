package Valid_Number

import "testing"

func Test_isNumber(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			"case",
			args{s: " "},
			false,
		},
		{
			"case",
			args{s: " 0.1"},
			true,
		},
		{
			"case",
			args{s: "abc"},
			false,
		},
		{
			"case",
			args{s: "1 a"},
			false,
		},
		{
			"case",
			args{s: "2e10"},
			true,
		},
		{
			"case",
			args{s: "-90e3"},
			true,
		},
		{
			"case",
			args{s: " 1e"},
			false,
		},
		{
			"case",
			args{s: "e3"},
			false,
		},
		{
			"case",
			args{s: " 6e-1"},
			true,
		},
		{
			"case",
			args{s: "99e2.5"},
			false,
		},
		{
			"case",
			args{s: "55.5e93"},
			true,
		},
		{
			"case",
			args{s: " --6"},
			false,
		},
		{
			"case",
			args{s: "+-4"},
			false,
		},
		{
			"case",
			args{s: ".1"},
			true,
		},
		{
			"case",
			args{s: "."},
			false,
		},
		{
			"case",
			args{s: "95a54e53"},
			false,
		},
		{
			"case",
			args{s: "3."},
			true,
		},
		{
			"case",
			args{s: "  2.8"},
			true,
		},
		{
			"case",
			args{s: ".-4"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isNumber(tt.args.s); got != tt.want {
				t.Errorf("isNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
