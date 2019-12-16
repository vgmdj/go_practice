package Maximum_Swap

import "testing"

func Test_maximumSwap(t *testing.T) {
	type args struct {
		num int
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
				98368,
			},
			98863,
		},
		{
			"case2",
			args{
				11345678432,
			},
			81345671432,
		},
		{
			"case3",
			args{
				9973,
			},
			9973,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumSwap(tt.args.num); got != tt.want {
				t.Errorf("maximumSwap() = %v, want %v", got, tt.want)
			}
		})
	}
}
