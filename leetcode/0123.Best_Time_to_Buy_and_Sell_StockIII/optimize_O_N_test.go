package Best_Time_to_Buy_and_Sell_StockII

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxProfitII(t *testing.T) {
	ast := assert.New(t)

	ast.Equal(0, maxProfitII([]int{8}))
	ast.Equal(5, maxProfitII([]int{1, 6}))
	ast.Equal(5, maxProfitII([]int{1, 2, 3, 4, 5, 6}))
	ast.Equal(9, maxProfitII([]int{7, 1, 6, 4, 8}))
	ast.Equal(15, maxProfitII([]int{8, 3, 6, 2, 8, 8, 8, 4, 2, 0, 7, 2, 9, 4, 9}))

}

