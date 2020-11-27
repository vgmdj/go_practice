package Maximum_Gap

import "testing"

func Test_maximumGap(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		// TODO: Add test cases.
		{
			"case",
			args{
				[]int{4, 2, 44, 45, 41, 37, 103, 543, 1324},
			},
			781,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := maximumGap(tt.args.nums); gotAns != tt.wantAns {
				t.Errorf("maximumGap() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
