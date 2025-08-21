package solutions

import (
	"fmt"
	"math"
)

// https://leetcode.com/problems/stone-game-ii
type StoneGame2Soln struct {
	dp    [][]int
	piles []int
}

func stoneGameII(piles []int) int {
	n := len(piles)
	soln := &StoneGame2Soln{
		dp:    make([][]int, n),
		piles: piles,
	}
	total := 0
	for i := range n {
		total += piles[i]
		soln.dp[i] = make([]int, n+1)
		for k := range n + 1 {
			soln.dp[i][k] = -1
		}
	}
	return soln.helper(0, 1, total)
}

func (s *StoneGame2Soln) helper(start, m, remaining int) int {
	n := len(s.piles)
	if start+2*m >= n {
		return remaining
	}

	if s.dp[start][m] != -1 {
		return s.dp[start][m]
	}

	otherPerson := math.MaxInt32 // minimize other person
	var left = remaining
	for i := 1; i <= 2*m; i++ { // Allowed range of jums
		left -= s.piles[start+i-1] // consumed by current person
		otherPerson = min(otherPerson, s.helper(start+i, max(i, m), left))
	}

	s.dp[start][m] = remaining - otherPerson
	return s.dp[start][m]
}

func StoneGame2() {
	fmt.Print(stoneGameII([]int{2, 7, 9, 4, 4}))
}
