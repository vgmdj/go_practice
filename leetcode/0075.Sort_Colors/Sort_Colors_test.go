package Sort_Colors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSort(t *testing.T) {
	ast := assert.New(t)

	case1 := []int{2, 0, 0, 1, 0, 1, 2, 2, 2, 0, 1, 2, 1, 0}
	sortColors(case1)
	ast.Equal([]int{0, 0, 0, 0, 0, 1, 1, 1, 1, 2, 2, 2, 2, 2}, case1)
}
