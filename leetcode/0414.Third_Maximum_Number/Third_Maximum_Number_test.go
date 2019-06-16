package Third_Maximum_Number

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThirdMaximumNumber(t *testing.T) {
	ast := assert.New(t)

	ast.Equal(2, thirdMax2([]int{1, 1,2}))
	ast.Equal(3, thirdMax2([]int{5, 3, 2, 1, 5, 6, 3, 2, 1}))
}
