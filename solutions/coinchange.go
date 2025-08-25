package solutions

import (
	"fmt"
	"slices"
)

// https://leetcode.com/problems/coin-change/description/

func coinChange(coins []int, amount int) int {
	change := make([]int, amount+1)
	slices.Sort(coins)
	for i := 1; i < len(change); i++ {
		change[i] = 9999999
		for j := 0; j < len(coins) && coins[j] <= i; j++ {
			change[i] = min(change[i], change[i-coins[j]]+1)
		}
	}
	if change[amount] == 9999999 {
		return -1
	}
	return change[amount]
}

func CoinChange() {
	out := coinChange([]int{2, 3, 6, 5}, 30)
	fmt.Printf("%v", out)
}
