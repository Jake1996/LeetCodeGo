package solutions

import "fmt"

func generateParenthesis(n int) []string {
	out := []string{}
	if n == 0 {
		return []string{""}
	}
	for i := 0; i < n; i++ {
		for _, parL := range generateParenthesis(i) {
			for _, parR := range generateParenthesis(n - i - 1) {
				out = append(out, fmt.Sprintf("(%v)%v", parL, parR))
			}
		}
	}
	return out
}

func GenerateParenthesis() {
	out := generateParenthesis(2)
	fmt.Printf("%v", out)
}
