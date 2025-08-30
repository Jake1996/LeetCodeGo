package solutions

import "fmt"

// https://leetcode.com/problems/maximal-square/
func maximalSquare(matrix [][]byte) int {
	ans := 0
	if len(matrix) == 0 {
		return ans
	}
	m := len(matrix)
	n := len(matrix[0])
	dp := make([][]int, m+1)
	onesRow := make([][]int, m+1)
	onesCol := make([][]int, m+1)
	for i := range len(dp) {
		dp[i] = make([]int, n+1)
		onesRow[i] = make([]int, n+1)
		onesCol[i] = make([]int, n+1)
	}

	// PreCompute OnesRow and OnesCol
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if matrix[i][j] == '1' {
				onesRow[i][j] = 1 + onesRow[i+1][j]
				onesCol[i][j] = 1 + onesCol[i][j+1]
			}
		}
	}

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if matrix[i][j] == '1' {
				dp[i][j] = 1 + min(onesCol[i][j]-1, onesRow[i][j]-1, dp[i+1][j+1])
				ans = max(ans, dp[i][j])
			}
		}
	}
	return ans * ans
}

func MaximalSquare() {
	fmt.Print(maximalSquare([][]byte{}))
}
