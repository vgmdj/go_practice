package Container_With_Most_Water

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxWater(t *testing.T) {
	ast := assert.New(t)

	ast.Equal(49, maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))

}
