package Single_Number_III

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleNumber(t *testing.T) {
	ast := assert.New(t)

	ast.Equal([]int{3, 5}, singleNumber([]int{1, 1, 2, 2, 3, 4, 4, 5}))

}
