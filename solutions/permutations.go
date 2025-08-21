package solutions

import "fmt"

type PermuteSoln struct {
	ans     [][]int
	nums    []int
	visited []bool
}

func permute(nums []int) [][]int {
	sol := &PermuteSoln{
		nums:    nums,
		visited: make([]bool, len(nums)),
	}
	sol.Recurse([]int{})
	return sol.ans
}

func (p *PermuteSoln) Recurse(cur []int) {
	if len(cur) == len(p.nums) {
		t := make([]int, len(cur))
		copy(t, cur)
		p.ans = append(p.ans, t)
		return
	}

	for k := 0; k < len(p.nums); k++ {
		if !p.visited[k] {
			cur = append(cur, p.nums[k])
			p.visited[k] = true
			p.Recurse(cur)
			p.visited[k] = false
			cur = cur[:len(cur)-1]
		}
	}
}

func Permutation() {
	out := permute([]int{1, 2, 3})
	fmt.Print(out)
}
