package Decrypt_String_from_Alphabet_to_Integer_Mapping

import "testing"

func Test_freqAlphabets(t *testing.T) {
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
			"case1",
			args{
				"10#11#12",
			},
			"jkab",
		},
		{
			"case2",
			args{
				"12345678910#11#12#13#14#15#16#17#18#19#20#21#22#23#24#25#26#",
			},
			"abcdefghijklmnopqrstuvwxyz",
		},
		{
			"case3",
			args{
				"10#",
			},
			"j",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := freqAlphabets(tt.args.s); got != tt.want {
				t.Errorf("freqAlphabets() = %v, want %v", got, tt.want)
			}
		})
	}
}