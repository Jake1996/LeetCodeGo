package solutions

import (
	"fmt"
	"math"
)

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/description/

func maxProfit(prices []int) int {
	minPrice := math.MaxInt
	profit := 0
	for _, price := range prices {
		minPrice = min(minPrice, price)
		profit = max(profit, price-minPrice)
	}
	return profit
}

func MaxProfit() {
	res := maxProfit([]int{1, 2, 3, 7, 3, 9})
	fmt.Printf("%v", res)
}
