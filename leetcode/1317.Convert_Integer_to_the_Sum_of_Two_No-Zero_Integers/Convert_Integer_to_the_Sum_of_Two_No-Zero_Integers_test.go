package Convert_Integer_to_the_Sum_of_Two_No_Zero_Integers

import (
	`reflect`
	`testing`
)

func Test_getNoZeroIntegers(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			"case",
			args{
				1057,
			},
			[]int{111,946},
		},
		{
			"case",
			args{
				10000,
			},
			[]int{1,9999},
		},
		{
			"case",
			args{
				11,
			},
			[]int{2,9},
		},
		{
			"case",
			args{
				69,
			},
			[]int{1,68},
		},
		{
			"case",
			args{
				11,
			},
			[]int{2,9},
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getNoZeroIntegers(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getNoZeroIntegers() = %v, want %v", got, tt.want)
			}
		})
	}
}