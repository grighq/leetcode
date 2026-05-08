package besttimetobuyandsellstock

func maxProfit(prices []int) int {
	profit := 0
	minPrice := prices[0]
	for _, currentPrice := range prices {
		if currentPrice < minPrice {
			minPrice = currentPrice
		} else if currentPrice-minPrice > profit {
			profit = currentPrice - minPrice
		}
	}

	return profit
}
