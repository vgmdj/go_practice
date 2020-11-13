package Freedom_Trail

import "testing"

func Test_findRotateSteps(t *testing.T) {
	type args struct {
		ring string
		key  string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			"case",
			args{
				ring: "godding",
				key:  "oi",
			},
			6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRotateSteps(tt.args.ring, tt.args.key); got != tt.want {
				t.Errorf("findRotateSteps() = %v, want %v", got, tt.want)
			}
		})
	}
}
