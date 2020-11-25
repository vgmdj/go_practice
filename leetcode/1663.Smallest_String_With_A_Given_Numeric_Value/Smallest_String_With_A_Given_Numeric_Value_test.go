package Smallest_String_With_A_Given_Numeric_Value

import "testing"

func Test_getSmallestString(t *testing.T) {
	type args struct {
		n int
		k int
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
				n: 5,
				k: 130,
			},
			"",
		},
		{
			"case",
			args{
				n: 24,
				k: 552,
			},
			"aadzzzzzzzzzzzzzzzzzzzzz",
		},
		{
			"case",
			args{
				n: 5,
				k: 73,
			},
			"aaszz",
		},
		{
			"case",
			args{
				n: 3,
				k: 27,
			},
			"aay",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSmallestString(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("getSmallestString() = %v, want %v", got, tt.want)
			}
		})
	}
}
