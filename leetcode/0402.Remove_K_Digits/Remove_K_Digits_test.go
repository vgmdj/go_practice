package Remove_K_Digits

import "testing"

func Test_removeKdigits(t *testing.T) {
	type args struct {
		num string
		k   int
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
				num: "10200",
				k:   1,
			},
			"200",
		},
		{
			"case",
			args{
				num: "736421",
				k:   2,
			},
			"3421",
		},
		{
			"case",
			args{
				num: "1234",
				k:   1,
			},
			"123",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeKdigits(tt.args.num, tt.args.k); got != tt.want {
				t.Errorf("removeKdigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
