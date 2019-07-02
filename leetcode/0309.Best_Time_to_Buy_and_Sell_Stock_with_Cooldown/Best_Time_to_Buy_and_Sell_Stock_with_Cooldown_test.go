package Best_Time_to_Buy_and_Sell_Stock_with_Cooldown

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxProfit(t *testing.T) {
	ast := assert.New(t)

	ast.Equal(3, maxProfit([]int{1, 2, 3, 0, 2}))

}
