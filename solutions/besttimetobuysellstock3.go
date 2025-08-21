package solutions

import "fmt"

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iii/description/

func maxProfit3(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	n := len(prices)
	maxProfit := 0
	min1 := prices[0]
	max1 := prices[n-1]
	pfw := make([]int, n)
	pbw := make([]int, n)
	for i := 1; i < n; i++ {
		min1 = min(min1, prices[i])
		max1 = max(max1, prices[n-i])
		pfw[i] = max(pfw[i-1], prices[i]-min1)
		pbw[n-i-1] = max(pbw[n-i], max1-prices[n-i])
	}
	for i := 1; i < n; i++ {
		maxProfit = max(maxProfit, pfw[i]+pbw[i])
	}
	return maxProfit
}

func MaxProfit3() {
	res := maxProfit3([]int{1, 2, 3, 4, 5})
	fmt.Printf("%v", res)
}
