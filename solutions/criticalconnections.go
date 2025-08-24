package solutions

import "fmt"

type CriticalConnSoln struct {
	adj   [][]int
	low   []int
	num   []int
	count int
	ans   [][]int
}

// https://leetcode.com/problems/critical-connections-in-a-network/
var ()

func (cc *CriticalConnSoln) dfs(u, par int) {
	cc.count++
	cc.num[u] = cc.count
	cc.low[u] = cc.count

	for _, v := range cc.adj[u] {
		if v == par {
			continue
		}
		if cc.num[v] == 0 { // Not visited
			cc.dfs(v, u)
			cc.low[u] = min(cc.low[u], cc.low[v])
			if cc.low[v] > cc.num[u] {
				cc.ans = append(cc.ans, []int{u, v})
			}
		} else {
			cc.low[u] = min(cc.low[u], cc.num[v])
		}
	}
}

func criticalConnections(n int, connections [][]int) [][]int {
	cc := &CriticalConnSoln{
		adj:   make([][]int, n),
		low:   make([]int, n),
		num:   make([]int, n),
		count: 0,
		ans:   make([][]int, 0, n),
	}
	for _, connection := range connections {
		u := connection[0]
		v := connection[1]
		cc.adj[u] = append(cc.adj[u], v)
		cc.adj[v] = append(cc.adj[v], u)
	}
	cc.dfs(0, -1)

	return cc.ans
}

func CriticalConn() {
	fmt.Print(criticalConnections(4, [][]int{{0, 1}, {1, 2}, {2, 0}, {1, 3}}))
}
