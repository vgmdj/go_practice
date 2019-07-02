package Best_Time_to_Buy_and_Sell_StockII

import "math"

func maxProfit3(prices []int) int {
	dp_i10, dp_i11 := 0, math.MinInt32
	dp_i20, dp_i21 := 0, math.MinInt32
	for _, price := range prices {
		dp_i20 = Max(dp_i20, dp_i21+price)
		dp_i21 = Max(dp_i21, dp_i10-price)
		dp_i10 = Max(dp_i10, dp_i11+price)
		dp_i11 = Max(dp_i11, -price)
	}
	return dp_i20

}
