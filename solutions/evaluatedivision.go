package solutions

import "container/list"

type edge struct {
	node string
	val  float64
}

// https://leetcode.com/problems/evaluate-division/
func CalcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	adjGraph := make(map[string][]edge)

	for i, eq := range equations {
		adjGraph[eq[0]] = append(adjGraph[eq[0]], edge{
			node: eq[1],
			val:  values[i],
		})
		adjGraph[eq[1]] = append(adjGraph[eq[1]], edge{
			node: eq[0],
			val:  1 / values[i],
		})
	}
	out := make([]float64, 0, len(queries))
	for _, query := range queries {
		out = append(out, dfs(adjGraph, query[0], query[1]))
	}
	return out
}

func dfs(graph map[string][]edge, src, dst string) float64 {
	queue := list.New()
	visited := make(map[string]bool)
	queue.PushBack(edge{
		node: src,
		val:  1,
	})
	for queue.Len() > 0 {
		el := queue.Front()
		queue.Remove(el)
		ed := el.Value.(edge)
		for _, v := range graph[ed.node] {
			if v.node == dst {
				return ed.val * v.val
			}
			if !visited[v.node] {
				queue.PushBack(edge{
					node: v.node,
					val:  v.val * ed.val,
				})
				visited[v.node] = true
			}
		}
	}
	return -1
}
