package solutions

import "fmt"

type AreConnectedSoln struct {
	union []int
	rank  []int
}

// https://leetcode.com/problems/graph-connectivity-with-threshold/description/
func areConnected(n int, threshold int, queries [][]int) []bool {
	con := &AreConnectedSoln{
		union: make([]int, n+1),
		rank:  make([]int, n+1),
	}

	for i := 1; i <= n; i++ {
		con.rank[i] = 1
		con.union[i] = i
	}

	for i := threshold + 1; i <= n; i++ {
		for j := 1; i*j <= n; j++ {
			con.join(i, i*j)
		}
	}
	out := make([]bool, len(queries))
	for i, query := range queries {
		if con.parent(query[0]) == con.parent(query[1]) {
			out[i] = true
		}
	}
	return out
}

func (c AreConnectedSoln) join(i, j int) {
	pi := c.parent(i)
	pj := c.parent(j)
	if c.rank[pi] > c.rank[pj] {
		c.union[pj] = pi
		c.rank[pi] += c.rank[pj]
	} else {
		c.union[pi] = pj
		c.rank[pj] += c.rank[pi]
	}
}

func (c AreConnectedSoln) parent(i int) int {
	if c.union[i] == i {
		return i
	}
	c.union[i] = c.parent(c.union[i])
	return c.union[i]
}

func AreConected() {
	out := areConnected(6, 2, [][]int{{1, 4}, {2, 5}, {3, 6}})
	fmt.Printf("%v", out)
}
