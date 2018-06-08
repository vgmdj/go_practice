package Best_Time_to_Buy_and_Sell_Stock

import "math"

func maxProfit(prices []int) int {
	minPrice, maxProfit := math.MaxInt32, 0

	for _, v := range prices {
		if minPrice > v {
			minPrice = v
		} else if v-minPrice > maxProfit {
			maxProfit = v - minPrice
		}
	}

	return maxProfit

}
