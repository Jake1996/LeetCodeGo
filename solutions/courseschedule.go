package solutions

import "fmt"

// https://leetcode.com/problems/course-schedule/
type CourseSchedule struct {
	graph    [][]int
	gvisited []bool
}

// returns true if cycle
func (c *CourseSchedule) dfs(i int, visited []bool) bool {
	visited[i] = true
	for _, v := range c.graph[i] {
		if c.gvisited[v] {
			continue
		}
		if visited[v] {
			return true
		}
		if c.dfs(v, visited) {
			return true
		}
	}
	visited[i] = false
	c.gvisited[i] = true
	return false
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	c := CourseSchedule{
		graph:    make([][]int, numCourses),
		gvisited: make([]bool, numCourses),
	}
	for _, prereq := range prerequisites {
		c.graph[prereq[0]] = append(c.graph[prereq[0]], prereq[1])
	}
	for i := range numCourses {
		if !c.gvisited[i] {
			if c.dfs(i, make([]bool, numCourses)) {
				return false
			}
		}
	}
	return true
}

func CanFinish() {
	fmt.Print(canFinish(2, [][]int{{1, 0}, {0, 1}}))
}
