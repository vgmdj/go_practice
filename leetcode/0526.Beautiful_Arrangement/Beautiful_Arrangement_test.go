package Beautiful_Arrangement

import (
	"testing"
)

func Test_countArrangement(t *testing.T) {
	type args struct {
		N int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test1",
			args{
				1,
			},
			1,
		},
		{
			"test2",
			args{
				2,
			},
			2,
		},
		{
			"test3",
			args{
				3,
			},
			3,
		},
		{
			"test4",
			args{
				4,
			},
			8,
		},
		{
			"test1",
			args{
				15,
			},
			24679,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countArrangement(tt.args.N); got != tt.want {
				t.Errorf("countArrangement() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isBeautiful(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test1",
			args{
				1,
				2,
			},
			true,
		},
		{
			"test2",
			args{
				6,
				3,
			},
			true,
		},
		{
			"test3",
			args{
				5,
				3,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isBeautiful(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("isBeautiful() = %v, want %v", got, tt.want)
			}
		})
	}
}
