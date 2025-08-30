package solutions

import "fmt"

type UniquePath2Soln struct {
	dp        [][]int
	obstacles [][]int
}

// https://leetcode.com/problems/unique-paths-ii/
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 {
		return 0
	}
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	soln := &UniquePath2Soln{
		dp:        make([][]int, m),
		obstacles: obstacleGrid,
	}
	for i := range m {
		soln.dp[i] = make([]int, n)
		for j := range n {
			soln.dp[i][j] = -1
		}
	}
	soln.dp[0][0] = 1
	return soln.recurse(m-1, n-1)
}

func (u *UniquePath2Soln) recurse(i, j int) int {
	if i < 0 || j < 0 || i >= len(u.dp) || j >= len(u.dp[0]) {
		return 0
	}
	if u.obstacles[i][j] == 1 {
		return 0
	}
	if u.dp[i][j] != -1 {
		return u.dp[i][j]
	}
	u.dp[i][j] = u.recurse(i-1, j) + u.recurse(i, j-1)
	return u.dp[i][j]
}

func UniquePaths2() {
	fmt.Print((uniquePathsWithObstacles([][]int{})))
}
