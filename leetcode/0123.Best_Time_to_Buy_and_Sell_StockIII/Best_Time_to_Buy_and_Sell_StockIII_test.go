package Best_Time_to_Buy_and_Sell_StockII

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxProfit(t *testing.T) {
	ast := assert.New(t)

	ast.Equal(0, maxProfit([]int{7}))
	ast.Equal(5, maxProfit([]int{1, 6}))
	ast.Equal(9, maxProfit([]int{7, 1, 6, 4, 8}))

}

func TestPeriodProfit(t *testing.T) {
	ast := assert.New(t)

	case1 := []int{7, 1, 5, 3, 6, 4}
	ast.Equal(periodMaxProfit(case1), 5)

	case2 := []int{7, 6, 4, 3, 1}
	ast.Equal(periodMaxProfit(case2), 0)

}