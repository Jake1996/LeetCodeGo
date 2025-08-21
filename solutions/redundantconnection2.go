package solutions

import "fmt"

// https://leetcode.com/problems/redundant-connection-ii/
type Redundant2Soln struct {
	parent []int
}

func (r *Redundant2Soln) Union(i, j int) bool {
	pi := r.Parent(i)
	pj := r.Parent(j)
	if pi == pj {
		return false
	}
	r.parent[pj] = pi
	return true
}

func (r *Redundant2Soln) Parent(i int) int {
	if r.parent[i] == i {
		return i
	}
	pi := r.Parent(r.parent[i])
	r.parent[i] = pi
	return pi
}

func findRedundantDirectedConnection(edges [][]int) []int {
	n := 0
	for _, edge := range edges {
		n = max(n, edge[0], edge[1])
	}
	parent := make(map[int]int)
	var ans1, ans2 []int
	r := &Redundant2Soln{
		parent: make([]int, n+1),
	}
	for i := range r.parent {
		r.parent[i] = i
	}
	for _, edge := range edges {
		if parent[edge[1]] != 0 {
			ans1 = edge                            // Two parent edge
			ans2 = []int{parent[edge[1]], edge[1]} // Previous edge incase cycle is detected
		}
		parent[edge[1]] = edge[0]
	}
	for _, edge := range edges {
		if ans1 != nil && edge[0] == ans1[0] && edge[1] == ans1[1] { // Skip parent edge
			continue
		}
		if u := r.Union(edge[0], edge[1]); !u { // Cycle detected
			if ans1 == nil {
				return edge
			} else {
				return ans2
			}
		}
	}
	return ans1 // No cycle detected
}

func FindRedundantConnection2() {
	out := findRedundantDirectedConnection([][]int{{2, 1}, {3, 1}, {4, 2}, {1, 4}})
	fmt.Print(out)
}
