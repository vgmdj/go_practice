package Coin_Change

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCoinChange(t *testing.T) {
	ast := assert.New(t)

	ast.Equal(3, coinChange([]int{1, 2, 5}, 11))
	ast.Equal(-1, coinChange([]int{2}, 3))
	ast.Equal(3, coinChange([]int{2, 3, 5}, 11))
	ast.Equal(3, coinChange([]int{2, 3, 5, 10}, 14))
	ast.Equal(5, coinChange([]int{2, 3, 5}, 21))
	ast.Equal(3, coinChange([]int{1, 2, 5}, 11))
	ast.Equal(3, coinChange([]int{1, 2, 5}, 11))
	ast.Equal(3, coinChange([]int{1, 2, 5}, 11))
	ast.Equal(3, coinChange([]int{1, 2, 5}, 11))
	ast.Equal(3, coinChange([]int{1, 2, 5}, 11))
	ast.Equal(3, coinChange([]int{1, 2, 5}, 11))
	ast.Equal(3, coinChange([]int{1, 2, 5}, 11))

}
