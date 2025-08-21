package solutions

import (
	"fmt"
	"slices"
)

// https://leetcode.com/problems/combination-sum/

func combinationSum(candidates []int, target int) [][]int {
	slices.Sort(candidates)

	dp := make([][][]int, target+1)
	dp[0] = make([][]int, 1)

	for _, c := range candidates {
		for i := c; i <= target; i++ {
			for _, ar := range dp[i-c] {
				newSL := make([]int, len(ar))
				copy(newSL, ar)
				newSL = append(newSL, c)
				dp[i] = append(dp[i], newSL)
			}
		}
	}

	return dp[target]
}

func CombinationSum() {
	out := combinationSum([]int{2, 3, 7}, 18)
	fmt.Printf("%v", out)
}
