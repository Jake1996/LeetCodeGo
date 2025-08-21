package solutions

import "fmt"

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/description/

func maxProfit2(prices []int) int {
	profit := 0
	for i := 1; i < len(prices); i++ {
		profit += max(0, prices[i]-prices[i-1])
	}

	return profit
}

func MaxProfit2() {
	res := maxProfit2([]int{1, 2, 3, 7, 3, 9})
	fmt.Printf("%v", res)
}
