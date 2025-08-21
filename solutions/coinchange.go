package solutions

import (
	"fmt"
)

// https://leetcode.com/problems/coin-change/description/

func coinChange(coins []int, amount int) int {
	change := make([]int, amount+1)
	for i := 1; i < len(change); i++ {
		change[i] = 9999999
		for j := 0; j < len(coins); j++ {
			if coins[j] <= i {
				change[i] = min(change[i], change[i-coins[j]]+1)
			}
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
