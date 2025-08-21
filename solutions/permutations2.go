package solutions

import (
	"fmt"
	"slices"
)

type PermuteUniqueSoln struct {
	ans     [][]int
	nums    []int
	visited []bool
}

func permuteUnique(nums []int) [][]int {
	slices.Sort(nums)
	sol := &PermuteUniqueSoln{
		nums:    nums,
		visited: make([]bool, len(nums)),
	}
	sol.RecurseUnique([]int{})
	return sol.ans
}

func (p *PermuteUniqueSoln) RecurseUnique(cur []int) {
	if len(cur) == len(p.nums) {
		t := make([]int, len(cur))
		copy(t, cur)
		p.ans = append(p.ans, t)
		return
	}

	for k := 0; k < len(p.nums); k++ {
		if k > 0 && p.nums[k] == p.nums[k-1] && !p.visited[k-1] {
			continue
		}
		if !p.visited[k] {
			cur = append(cur, p.nums[k])
			p.visited[k] = true
			p.RecurseUnique(cur)
			p.visited[k] = false
			cur = cur[:len(cur)-1]
		}
	}
}

func PermutationUnique() {
	out := permuteUnique([]int{1, 1, 3})
	fmt.Print(out)
}
