package solutions

import "fmt"

// https://leetcode.com/problems/wildcard-matching/
func isMatch(s string, p string) bool {
	var i, j, match int
	startI := -1
	for i < len(s) {
		if j < len(p) && (s[i] == p[j] || p[j] == '?') {
			i++
			j++
		} else if j < len(p) && p[j] == '*' {
			startI = j
			match = i // last matched position
			j++
		} else if startI != -1 { // Assume the characters are part of a *
			i = match + 1
			j = startI + 1
			match++
		} else {
			return false
		}
	}
	for j < len(p) && p[j] == '*' {
		j++
	}
	return j == len(p)
}

func IsMatch() {
	fmt.Print(isMatch("abcde", "a?c*"))
}
