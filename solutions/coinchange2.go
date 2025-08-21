package solutions

import (
	"fmt"
)

// https://leetcode.com/problems/coin-change-ii/description/

func change(amount int, coins []int) int {

	dp := make([]int, amount+1)

	dp[0] = 1

	// Iterate through each coin denomination.
	for _, coin := range coins {
		// For each coin, iterate from the coin's value up to the total amount.
		// This ensures we are only considering the current coin and coins before it.
		for i := coin; i <= amount; i++ {
			// The number of combinations for amount `i` is the sum of:
			// 1. Combinations without using the current `coin` (already in `dp[i]`).
			// 2. Combinations by using the current `coin`. This is equal to the
			//    number of combinations for the remaining amount `i-coin`.
			dp[i] += dp[i-coin]
		}
	}

	// The final answer is the number of combinations for the total amount.
	return dp[amount]
}

func Change() {
	out := change(500, []int{2, 7, 13})
	fmt.Printf("%v", out)
}
