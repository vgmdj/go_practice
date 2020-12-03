package Maximum_Repeating_Substring

import "testing"

func Test_maxRepeating(t *testing.T) {
	type args struct {
		sequence string
		word     string
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
				sequence: "bbbbbb",
				word:     "bb",
			},
			3,
		},
		{
			"case",
			args{
				sequence: "ababc",
				word:     "ba",
			},
			1,
		},
		{
			"case",
			args{
				sequence: "cabcabc",
				word:     "ab",
			},
			1,
		},
		{
			"case",
			args{
				sequence: "cababc",
				word:     "ab",
			},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxRepeating(tt.args.sequence, tt.args.word); got != tt.want {
				t.Errorf("maxRepeating() = %v, want %v", got, tt.want)
			}
		})
	}
}
