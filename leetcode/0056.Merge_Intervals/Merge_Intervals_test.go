package Merge_Intervals

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	ast := assert.New(t)

	ast.Equal([]Interval{
		{Start: 0, End: 4},
	}, merge([]Interval{
		{1, 4},
		{0, 4},
	}))

}
