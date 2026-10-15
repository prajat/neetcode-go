package main

func MaxProfit(prices []int) int {
	maxProfit := 0
	minPrice := prices[0]
	for _, price := range prices[1:] {
		// Try selling today
		profit := price - minPrice
		maxProfit = max(profit, maxProfit)
		// Try today's price as minPrice contender
		minPrice = min(minPrice, price)
	}
	return maxProfit
}
