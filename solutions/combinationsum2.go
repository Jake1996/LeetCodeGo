package solutions

import (
	"fmt"
	"slices"
)

// https://leetcode.com/problems/combination-sum-ii/description/
type CombinationSumSoln struct {
	ans [][]int
}

func combinationSum2(candidates []int, target int) [][]int {
	out := CombinationSumSoln{
		ans: [][]int{},
	}
	slices.Sort(candidates)
	out.helpercombsum2(candidates, 0, target, []int{})
	return out.ans
}

func (c *CombinationSumSoln) helpercombsum2(candidates []int, start, target int, cur []int) {
	if target == 0 {
		c.ans = append(c.ans, cur)
	}
	if start >= len(candidates) || target < 0 {
		return
	}

	for i := start; i < len(candidates); i++ {
		if i > start && candidates[i] == candidates[i-1] {
			continue
		}
		if candidates[i] <= target {
			t := make([]int, len(cur))
			copy(t, cur)
			t = append(t, candidates[i])
			c.helpercombsum2(candidates, i+1, target-candidates[i], t)
		}
	}
}

func CombinationSum2() {
	out := combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8)
	fmt.Printf("%v", out)
}
