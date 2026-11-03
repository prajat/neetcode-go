package main

/*
You are given an integer array coins representing coins of different denominations (e.g. 1 dollar, 5 dollars, etc) and an integer amount representing a target amount of money.
Return the fewest number of coins that you need to make up the exact target amount. If it is impossible to make up the amount, return -1.
You may assume that you have an unlimited number of each coin.
*/

// using recursion only - TLE
func coinChange(coins []int, amount int) int {
	var dfs func(int) int
	dfs = func(remaining int) int {
		// exact amount formed
		if remaining == 0 {
			return 0
		}

		if remaining < 0 {
			return amount + 1
		}

		minCoins := amount + 1

		for _, coin := range coins {
			coinsNeeded := dfs(remaining - coin)

			if coinsNeeded+1 < minCoins {
				minCoins = coinsNeeded + 1
			}
		}

		return minCoins
	}

	result := dfs(amount)

	// for impossible case
	if result == amount+1 {
		return -1
	}

	return result
}

// recursion + memoisation
func coinChangeMemo(coins []int, amount int) int {
	cache := make(map[int]int, 0)

	var dfs func(int) int
	dfs = func(remaining int) int {
		// exact amount formed
		if remaining == 0 {
			return 0
		}

		if remaining < 0 {
			return amount + 1
		}

		if result, exists := cache[remaining]; exists {
			return result
		}

		minCoins := amount + 1
		for _, coin := range coins {
			coinsNeeded := dfs(remaining - coin)

			if coinsNeeded+1 < minCoins {
				minCoins = coinsNeeded + 1
			}
		}

		cache[remaining] = minCoins

		return cache[remaining]
	}

	result := dfs(amount)

	// for impossible case
	if result == amount+1 {
		return -1
	}

	return result
}
