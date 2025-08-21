package solutions

import "fmt"

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iv/description/

func maxProfit4(k int, prices []int) int {
	n := len(prices)
	if len(prices) < 2 {
		return 0
	}
	profits := make([][]int, k+1)
	for i := range profits {
		profits[i] = make([]int, n+1) // For each row, create 'y' columns
	}
	for i := 1; i <= k; i++ {
		curProfit := -1 * prices[0]
		for j := 1; j <= n; j++ {
			// Max Profit after completing i-1 transactions on day j-1 and buying on j-1
			curProfit = max(curProfit, profits[i-1][j-1]-prices[j-1])
			// Profit after making i transactions on day j
			// profits[i][j-1] - i transactions uptill day j-1
			//  curProfit+prices[j-1] - curProfit and sell at prices [j-1]
			profits[i][j] = max(profits[i][j-1], curProfit+prices[j-1])
		}
	}
	return profits[k][n]
}

func MaxProfit4() {
	res := maxProfit4(2, []int{3, 2, 6, 5, 0, 3})
	fmt.Printf("%v", res)
}
