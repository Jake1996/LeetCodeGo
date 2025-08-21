package solutions

import (
	"container/list"
	"fmt"
)

// https://leetcode.com/problems/course-schedule-iv/description/
func checkIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool {
	out := make([]bool, len(queries))
	if len(prerequisites) == 0 || len(queries) == 0 {
		return out
	}
	visited := make([]int, numCourses)
	leaves := list.New()
	preMap := make([]map[int]bool, numCourses)
	for i := range numCourses {
		preMap[i] = make(map[int]bool)
	}
	graph := make(map[int][]int)
	for _, prereq := range prerequisites {
		visited[prereq[1]]++
		graph[prereq[0]] = append(graph[prereq[0]], prereq[1])
	}
	for i, v := range visited {
		if v == 0 {
			leaves.PushBack(i)
		}
	}
	for leaves.Len() > 0 {
		leaf := leaves.Front().Value.(int)
		leaves.Remove(leaves.Front())
		nodes := graph[leaf]
		for _, node := range nodes {
			visited[node]--
			hs := preMap[node]
			hs[leaf] = true
			for k := range preMap[leaf] {
				hs[k] = true
			}
			if visited[node] <= 0 {
				leaves.PushBack(node)
			}
		}
	}

	for i, query := range queries {
		nodeMap := preMap[query[1]]
		if nodeMap[query[0]] {
			out[i] = true
		}
	}
	return out
}

func CheckIfPrerequisite() {
	out := checkIfPrerequisite(3, [][]int{{1, 2}, {1, 0}, {2, 0}}, [][]int{{1, 0}, {1, 2}})
	fmt.Printf("%v", out)
}
