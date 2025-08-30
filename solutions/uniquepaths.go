package solutions

import "fmt"

type UniquePathSoln struct {
	dp [][]int
}

// https://leetcode.com/problems/unique-paths/
func uniquePaths(m int, n int) int {
	soln := &UniquePathSoln{
		dp: make([][]int, m),
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

func (u *UniquePathSoln) recurse(i, j int) int {
	if i < 0 || j < 0 || i >= len(u.dp) || j >= len(u.dp[0]) {
		return 0
	}
	if u.dp[i][j] != -1 {
		return u.dp[i][j]
	}
	u.dp[i][j] = u.recurse(i-1, j) + u.recurse(i, j-1)
	return u.dp[i][j]
}

func UniquePaths() {
	fmt.Print(uniquePaths(3, 2))
}
