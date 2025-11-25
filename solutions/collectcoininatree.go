package solutions

// https://leetcode.com/problems/collect-coins-in-a-tree/
func CollectTheCoins(coins []int, edges [][]int) int {
	inDegree := make([]int, len(coins))
	graph := make(map[int][]int)
	for _, edge := range edges {
		inDegree[edge[0]]++
		inDegree[edge[1]]++
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}
	queue := make([]int, 0, len(coins))
	for i, in := range inDegree {
		if in == 1 && coins[i] == 0 {
			queue = append(queue, i)
		}
	}
	// Remove all leaves with no coins.
	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		inDegree[u]--
		for _, v := range graph[u] {
			inDegree[v]--
			if inDegree[v] == 1 && coins[v] == 0 {
				queue = append(queue, v)
			}
		}
		delete(graph, u)
	}
	// Remove leaves twice as we can directly pick them
	for range 2 {
		for i, in := range inDegree {
			if in == 1 {
				queue = append(queue, i)
			}
		}
		for _, u := range queue {
			queue = queue[1:]
			inDegree[u]--
			for _, v := range graph[u] {
				inDegree[v]--
			}
			delete(graph, u)
		}
	}

	remaining := 0
	for _, edge := range edges {
		if inDegree[edge[0]] > 0 && inDegree[edge[1]] > 0 {
			// Count all remaining edges.
			remaining++
		}
	}
	// Twice remaining as we need to come back to start position.
	return remaining * 2
}
