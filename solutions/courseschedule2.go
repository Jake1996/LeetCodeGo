package solutions

import (
	"container/list"
	"fmt"
)

// https://leetcode.com/problems/course-schedule-ii/
func findOrder(numCourses int, prerequisites [][]int) []int {
	inflow := make([]int, numCourses)
	graph := make(map[int][]int)

	out := make([]int, 0)

	q := list.New()

	for _, pre := range prerequisites {
		graph[pre[1]] = append(graph[pre[1]], pre[0])
		inflow[pre[0]]++
	}

	for i := range numCourses {
		if inflow[i] == 0 {
			q.PushBack(i)
		}
	}

	for q.Len() > 0 {
		e := q.Front()
		q.Remove(e)
		u := e.Value.(int)
		out = append(out, u)
		for _, v := range graph[u] {
			inflow[v]--
			if inflow[v] == 0 {
				q.PushBack(v)
			}
		}
	}
	if len(out) < numCourses {
		return []int{}
	}
	return out
}

func CourseSchedule2() {
	fmt.Print(findOrder(3, [][]int{}))
}
