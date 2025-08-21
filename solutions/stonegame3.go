package solutions

import (
	"fmt"
	"math"
)

// https://leetcode.com/problems/stone-game-iii/
type StoneGame3Soln struct {
	dp    []int
	piles []int
}

func stoneGameIII(piles []int) string {
	n := len(piles)
	soln := &StoneGame3Soln{
		dp:    make([]int, n+1),
		piles: piles,
	}
	total := 0
	for i := range n {
		total += piles[i]
		soln.dp[i] = -1
	}
	alice := soln.helper(0, total)
	if total-alice < alice {
		return "Alice"
	} else if total-alice > alice {
		return "Bob"
	}
	return "Tie"
}

func (s *StoneGame3Soln) helper(start, remaining int) int {
	n := len(s.piles)

	if s.dp[start] != -1 {
		return s.dp[start]
	}

	currentPerson := s.piles[start]
	otherPerson := math.MaxInt32 // minimize other person
	if start+1 <= n {
		otherPerson = min(otherPerson, s.helper(start+1, remaining-currentPerson))
	}
	if start+2 <= n {
		currentPerson += s.piles[start+1]
		otherPerson = min(otherPerson, s.helper(start+2, remaining-currentPerson))
	}
	if start+3 <= n {
		currentPerson += s.piles[start+2]
		otherPerson = min(otherPerson, s.helper(start+3, remaining-currentPerson))

	}
	if otherPerson != math.MaxInt32 {
		s.dp[start] = remaining - otherPerson

	}
	return s.dp[start]
}

func StoneGame3() {
	fmt.Print(stoneGameIII([]int{1, 2, 3, 7}))
}
