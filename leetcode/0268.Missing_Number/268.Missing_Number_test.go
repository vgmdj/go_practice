package Missing_Number

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMissing(t *testing.T) {
	ast := assert.New(t)

	case1 := []int{0, 1, 2, 3, 4, 5, 7}
	ast.Equal(missingNumber(case1), 6)

	case2 := []int{0, 9, 5, 4, 2, 3, 6, 7, 1}
	ast.Equal(missingNumber(case2), 8)

	case3 := []int{1}
	ast.Equal(missingNumber(case3), 0)

	case4 := []int{0, 1}
	ast.Equal(missingNumber(case4), 2)

}
