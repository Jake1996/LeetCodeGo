package solutions

import (
	"fmt"
	"slices"
)

type CombinationSum3Soln struct {
	ans        [][]int
	candidates []int
}

func combinationSum3(k int, target int) [][]int {
	out := CombinationSum3Soln{
		ans:        [][]int{},
		candidates: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
	}
	slices.Sort(out.candidates)
	out.helpercombsum3(0, 0, k, target, []int{})
	return out.ans
}

func (c *CombinationSum3Soln) helpercombsum3(start, j, k, target int, cur []int) {
	if target == 0 && j == k {
		c.ans = append(c.ans, cur)
	}
	if start >= len(c.candidates) || target < 0 {
		return
	}

	for i := start; i < len(c.candidates); i++ {
		if c.candidates[i] <= target {
			t := make([]int, len(cur))
			copy(t, cur)
			t = append(t, c.candidates[i])
			c.helpercombsum3(i+1, j+1, k, target-c.candidates[i], t)
		}
	}
}

func CombinationSum3() {
	out := combinationSum3(3, 8)
	fmt.Printf("%v", out)
}
