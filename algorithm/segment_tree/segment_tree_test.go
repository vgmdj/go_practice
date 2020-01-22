package segment_tree

import (
	"testing"
)

func TestSegmentTree_Build1(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want []int
	}{
		// TODO: Add test cases.
		{
			"case",
			[]int{1, 2, 3, 4},
			[]int{10, 3, 7, 1, 2, 3, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			st := NewSegmentTree(tt.data)
			st.Build()

			if len(st.sum) != len(tt.want) {
				t.Errorf("excepted count %d , but got count %d\n", len(tt.want), len(st.sum))
				return
			}

			for i := range tt.want {
				if tt.want[i] != st.sum[i] {
					t.Errorf("excepted sum num %d , but got %d\n", tt.want[i], st.sum[i])
				}
			}

		})
	}
}
