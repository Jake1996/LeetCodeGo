package solutions

import "fmt"

func wordBreak(s string, wordDict []string) bool {
	hs := make(map[string]bool)

	for _, word := range wordDict {
		hs[word] = true
	}
	dp := make([]bool, len(s))
	for i := 0; i < len(s); i++ {
		for j := i - 1; !dp[i] && j >= 0; j-- {
			sub := s[j : i+1]
			if hs[sub] {
				if j > 0 {
					dp[i] = dp[i] || dp[j-1]
				} else {
					dp[i] = true
				}
			}
		}
	}

	return dp[len(s)-1]
}

func WordBreak() {
	fmt.Print(wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
}
