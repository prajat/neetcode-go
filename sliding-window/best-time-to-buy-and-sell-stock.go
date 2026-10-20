package main

/*You are given an integer array prices where prices[i] is the price of NeetCoin on the ith day.
You may choose a single day to buy one NeetCoin and choose a different day in the future to sell it.
Return the maximum profit you can achieve. You may choose to not make any transactions, in which case the profit would be 0. */

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
