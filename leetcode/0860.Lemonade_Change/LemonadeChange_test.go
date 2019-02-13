package Lemonade_Change

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLemonade(t *testing.T) {
	ast := assert.New(t)

	ast.Equal(true, lemonadeChange([]int{5, 5, 5, 10, 20}))
	ast.Equal(true, lemonadeChange([]int{5, 5, 10}))
	ast.Equal(false, lemonadeChange([]int{5, 5, 10, 10, 20}))
	ast.Equal(false, lemonadeChange([]int{10, 10}))

}
