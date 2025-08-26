package solutions

import "fmt"

// https://leetcode.com/problems/wildcard-matching/
func isMatch(s string, p string) bool {
	var i, j, match int
	lastStar := -1
	for i < len(s) {
		if j < len(p) && (s[i] == p[j] || p[j] == '?') {
			i++
			j++
		} else if j < len(p) && p[j] == '*' {
			lastStar = j
			match = i // last matched position
			j++
		} else if lastStar != -1 { // Assume the characters are part of a *
			i = match + 1
			j = lastStar + 1
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
