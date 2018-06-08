package Majority_Element

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestME(t *testing.T) {
	ast := assert.New(t)

	case1 := []int{1, 1, 1, 1, 1, 3, 4, 5}
	ast.Equal(majorityElement(case1), 1)

	case2 := []int{1, 2, 3, 4, 1, 1, 1, 1}
	ast.Equal(majorityElement(case2), 1)

	case3 := []int{3, 3, 4}
	ast.Equal(majorityElement(case3), 3)

	case4 := []int{2, 2, 1, 1, 1, 2, 2}
	ast.Equal(majorityElement(case4), 2)

}
