package solutions

import "fmt"

func wordBreak2(s string, wordDict []string) []string {
	hs := make(map[string]bool)

	for _, word := range wordDict {
		hs[word] = true
	}

	dp := make([][]string, len(s))
	for i := 0; i < len(s); i++ {
		for j := 0; j <= i; j++ {
			sub := s[j : i+1]
			if hs[sub] {
				if j > 0 {
					for _, js := range dp[j-1] {
						dp[i] = append(dp[i], fmt.Sprintf("%v %v", js, sub))
					}
				} else {
					dp[i] = append(dp[i], sub)
				}
			}
		}
	}

	return dp[len(s)-1]
}

func WordBreak2() {
	fmt.Print(wordBreak2("catsanddog", []string{"cats", "dog", "sand", "and", "cat"}))
}
