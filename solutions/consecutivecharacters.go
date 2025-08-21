package solutions

import "fmt"

// https://leetcode.com/problems/consecutive-characters/
func maxPower(s string) int {
	var prev rune
	var m, i int
	for _, c := range s {
		if c == prev {
			i++
		} else {
			i = 1
		}
		prev = c
		m = max(m, i)
	}
	return m
}

func MaxPower() {
	out := maxPower("abbbbcddee")
	fmt.Printf("%v", out)
}
