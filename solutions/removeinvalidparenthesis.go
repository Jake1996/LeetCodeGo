package solutions

import (
	"fmt"
)

type InvalidParenthesisSoln struct {
	ans    []string
	ansMap map[string]bool
}

func removeInvalidParentheses(s string) []string {
	sol := &InvalidParenthesisSoln{
		ansMap: make(map[string]bool),
	}
	var open, close int
	for _, c := range s {
		if c == '(' {
			open++
		} else if c == ')' {
			if open == 0 {
				close++
			}
			if open > 0 {
				open--
			}
		}
	}
	sol.traverse(s, 0, 0, 0, open, close, "")
	for k := range sol.ansMap {
		sol.ans = append(sol.ans, k)
	}
	return sol.ans
}

func (ip *InvalidParenthesisSoln) traverse(s string, cur, o, c, ro, rc int, sb string) {
	if len(s) == cur && ro == 0 && rc == 0 {
		ip.ansMap[sb] = true
		return
	}
	if cur >= len(s) {
		return
	}
	ch := s[cur : cur+1]
	if ch == "(" && ro > 0 { // skip open
		ip.traverse(s, cur+1, o, c, ro-1, rc, sb)
	} else if ch == ")" && rc > 0 { // skip close
		ip.traverse(s, cur+1, o, c, ro, rc-1, sb)
	}
	sb = fmt.Sprintf("%v%v", sb, ch) // Include character
	if ch != "(" && ch != ")" {
		ip.traverse(s, cur+1, o, c, ro, rc, sb)
	}
	if ch == "(" {
		ip.traverse(s, cur+1, o+1, c, ro, rc, sb)
	}
	if ch == ")" && o > c {
		ip.traverse(s, cur+1, o, c+1, ro, rc, sb)
	}
}

func RemoveInvalidParenthesis() {
	fmt.Print(removeInvalidParentheses(")("))
}
