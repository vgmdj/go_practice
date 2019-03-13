package Teemo_Attacking

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDuration(t *testing.T) {
	ast := assert.New(t)

	ast.Equal(4, findPoisonedDuration([]int{1, 4}, 2))
	ast.Equal(3, findPoisonedDuration([]int{1, 2}, 2))
}
