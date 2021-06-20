package bytedance

import "testing"

func TestDiff(t *testing.T) {
	type args struct {
		data []int
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
				[]int{1,2,4,6,3,2,2},
			},
			5,
		},
		{
			"case2",
			args{
				[]int{2,2,2,3,4,5,5,6,6,5,4,4,1},
			},
			6,
		},
		{
			"case3",
			args{
				[]int{1,1,2,2,2,1},
			},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Diff(tt.args.data); got != tt.want {
				t.Errorf("Diff() = %v, want %v", got, tt.want)
			}
		})
	}
}
