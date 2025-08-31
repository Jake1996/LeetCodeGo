package solutions

import (
	"fmt"
	"slices"
)

// https://leetcode.com/problems/burst-balloons/
func maxCoins(nums []int) int {
	n := len(nums)
	arr := make([]int, n+2)

	// Add Virtual characters
	arr[0] = 1
	arr[len(arr)-1] = 1
	arr = slices.Replace(arr, 1, 1+n, nums...)

	dp := make([][]int, n+2)
	for i := range len(dp) {
		dp[i] = make([]int, n+2)
	}

	for length := 1; length <= n; length++ {
		for left := 1; left <= n-length+1; left++ {
			right := left + length - 1
			for k := left; k <= right; k++ {
				// Values from Left to k-1 and k+1 to right are already picked
				curskill := arr[left-1] * arr[k] * arr[right+1]
				dp[left][right] = max(dp[left][right], dp[left][k-1]+curskill+dp[k+1][right])
			}
		}
	}

	return dp[1][n]
}

func MaxCoins() {
	fmt.Print(maxCoins([]int{3, 1, 5, 8}))
}
