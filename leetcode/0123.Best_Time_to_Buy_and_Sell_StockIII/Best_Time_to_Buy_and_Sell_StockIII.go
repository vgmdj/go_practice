package Best_Time_to_Buy_and_Sell_StockII

func maxProfit(prices []int) int {
	max := 0
	for i := 1; i < len(prices)-1; i++ {
		if prices[i] <= prices[i-1] || prices[i] < prices[i+1]{
			continue
		}

		profit := periodMaxProfit(prices[:i+1]) + periodMaxProfit(prices[i+1:])
		if max < profit {
			max = profit
		}
	}

	profit := periodMaxProfit(prices)
	if max < profit {
		max = profit
	}

	return max
}

func periodMaxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	min, profit := prices[0], 0

	for _, v := range prices {
		if v < min {
			min = v
		}

		if profit < v-min {
			profit = v - min
		}
	}

	return profit

}
