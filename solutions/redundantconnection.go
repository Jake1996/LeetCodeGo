package solutions

import "fmt"

// https://leetcode.com/problems/redundant-connection/description/
type RedundantSoln struct {
	parent []int
	rank   []int
}

func (r *RedundantSoln) Union(i, j int) bool {
	pi := r.Parent(i)
	pj := r.Parent(j)
	if pi == pj {
		return false
	}
	if r.rank[pi] > r.rank[pj] {
		r.parent[pj] = pi
		r.rank[pi] += r.rank[pj]
	} else {
		r.parent[pi] = pj
		r.rank[pj] += r.rank[pi]

	}
	return true
}

func (r *RedundantSoln) Parent(i int) int {
	if r.parent[i] == i {
		return i
	}
	r.parent[i] = r.Parent(r.parent[i])
	return r.parent[i]
}

func findRedundantConnection(edges [][]int) []int {
	n := 0
	for _, edge := range edges {
		n = max(n, edge[0], edge[1])
	}

	r := &RedundantSoln{
		parent: make([]int, n+1),
		rank:   make([]int, n+1),
	}
	for i := range r.parent {
		r.parent[i] = i
		r.rank[i] = 1
	}
	for _, edge := range edges {
		if u := r.Union(edge[0], edge[1]); !u {
			return edge
		}
	}
	return []int{}
}

func FindRedundantConnection() {
	out := findRedundantConnection([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}})
	fmt.Print(out)
}
