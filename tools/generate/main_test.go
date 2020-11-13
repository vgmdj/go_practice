package main

import "testing"

func Test_dealFileName(t *testing.T) {
	type args struct {
		str string
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
				"23. Merge k Sorted Lists",
			},
			"0023.Merge_k_Sorted_Lists",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dealFileName(tt.args.str); got != tt.want {
				t.Errorf("dealFileName() = %v, want %v", got, tt.want)
			}
		})
	}
}
