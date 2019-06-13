package Search_a_2D_Matrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch2DMatrix(t *testing.T) {
	ast := assert.New(t)

	ast.Equal(false, searchMatrix([][]int{{1, 1}}, 2))
}
