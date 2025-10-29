package solutions

import "fmt"

// https://leetcode.com/problems/climbing-stairs/description/
var memo = make(map[int]int)

func climbStairs(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	} else if memo[n] == 0 {
		memo[n] = climbStairs(n-1) + climbStairs(n-2)
	}
	return memo[n]
}

func ClimbStairs() {
	out := climbStairs(5)
	fmt.Printf("%v", out)
}
