package Unique_Number_of_Occurrences

import "testing"

func Test_uniqueOccurrences(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			"case",
			args{
				arr: []int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0},
			},
			true,
		},
		{
			"case",
			args{
				arr: []int{1, 2},
			},
			false,
		},
		{
			"case",
			args{
				arr: []int{1, 2, 2},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniqueOccurrences(tt.args.arr); got != tt.want {
				t.Errorf("uniqueOccurrences() = %v, want %v", got, tt.want)
			}
		})
	}
}
