package Simplify_Path

import "testing"

func Test_simplifyPath(t *testing.T) {
	type args struct {
		path string
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
				"/../",
			},
			"/",
		},
		{
			"case",
			args{
				"/home/..////../a/",
			},
			"/a",
		},
		{
			"case",
			args{
				"/home/",
			},
			"/home",
		},
		{
			"case",
			args{
				"/home////b/a/c/../../d",
			},
			"/home/b/d",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := simplifyPath(tt.args.path); got != tt.want {
				t.Errorf("simplifyPath() = %v, want %v", got, tt.want)
			}
		})
	}
}