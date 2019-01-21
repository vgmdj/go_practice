package Best_Time_to_Buy_and_Sell_StockII

func maxProfitII(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	pf := make([]int, len(prices))
	pl := make([]int, len(prices))

	minPrice := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		}
		pf[i] = Max(prices[i]-minPrice, pf[i-1])
	}

	maxPrice := prices[len(prices)-1]
	for i := len(prices) - 2; i >= 0; i-- {
		if prices[i] > maxPrice {
			maxPrice = prices[i]
		}

		pl[i] = Max(maxPrice-prices[i], pl[i+1])
	}

	profit := 0
	for i := 0; i < len(prices); i++ {
		profit = Max(pf[i]+pl[i], profit)
	}

	return profit

}

func Max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
