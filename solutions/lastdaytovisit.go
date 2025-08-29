package solutions

import "fmt"

// https://leetcode.com/problems/last-day-where-you-can-still-cross/description/
type LatestDaySoln struct {
	dsu     []int
	visited [][]bool
}

func (l *LatestDaySoln) Union(i, j int) {
	pi := l.Parent(i)
	pj := l.Parent(j)
	l.dsu[pj] = pi
}

func (l *LatestDaySoln) Parent(i int) int {
	if l.dsu[i] == i {
		return i
	}
	l.dsu[i] = l.Parent(l.dsu[i])
	return l.dsu[i]
}

func latestDayToCross(row int, col int, cells [][]int) int {
	soln := &LatestDaySoln{
		dsu:     make([]int, row*col+2),
		visited: make([][]bool, row),
	}
	for i := range len(soln.dsu) {
		soln.dsu[i] = i
	}
	for i := range len(soln.visited) {
		soln.visited[i] = make([]bool, col)
	}
	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	virtualStart := row * col
	virtualEnd := row*col + 1
	for i := len(cells) - 1; i >= 0; i-- {
		r := cells[i][0] - 1
		c := cells[i][1] - 1
		soln.visited[r][c] = true
		curr := r*col + c
		for _, d := range dirs {
			nr := r + d[0]
			nc := c + d[1]
			if nr >= row || nr < 0 || nc >= col || nc < 0 {
				continue
			}
			s := nr*col + nc
			if soln.visited[nr][nc] {
				soln.Union(curr, s)
			}
		}
		if r == 0 {
			soln.Union(virtualStart, curr)
		}
		if r == row-1 {
			soln.Union(virtualEnd, curr)
		}
		if soln.Parent(virtualEnd) == soln.Parent(virtualStart) {
			return i
		}
	}
	return 0
}

func LatestDayToCross() {
	fmt.Print(latestDayToCross(2, 2, [][]int{{1, 1}, {2, 1}, {1, 2}, {2, 2}}))
}
