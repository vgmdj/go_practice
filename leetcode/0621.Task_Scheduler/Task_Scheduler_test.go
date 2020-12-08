package Task_Scheduler

import "testing"

func Test_leastInterval(t *testing.T) {
	type args struct {
		tasks []byte
		n     int
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
				[]byte{'A', 'A', 'A', 'A', 'A', 'A', 'B', 'C', 'D', 'E', 'F', 'G'},
				2,
			},
			16,
		},
		{
			"case",
			args{
				[]byte{'A', 'A', 'A', 'B', 'B', 'B'},
				1,
			},
			6,
		},
		{
			"case",
			args{
				[]byte{'A', 'A', 'A', 'B', 'B', 'B'},
				2,
			},
			8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := leastInterval(tt.args.tasks, tt.args.n); got != tt.want {
				t.Errorf("leastInterval() = %v, want %v", got, tt.want)
			}
		})
	}
}
