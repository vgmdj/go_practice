package Best_Time_to_Buy_and_Sell_Stock

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBestProfit(t *testing.T) {
	ast := assert.New(t)

	case1 := []int{7, 1, 5, 3, 6, 4}
	ast.Equal(maxProfit(case1), 5)

	case2 := []int{7, 6, 4, 3, 1}
	ast.Equal(maxProfit(case2), 0)

}
