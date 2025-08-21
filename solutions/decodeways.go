package solutions

import (
	"fmt"
	"strconv"
)

// https://leetcode.com/problems/decode-ways/description/
func numDecodings(s string) int {
	dp := make([]int, len(s))

	if s[0] == '0' {
		return 0
	}
	dp[0] = 1
	for i := 1; i < len(s); i++ {
		if s[i] != '0' {
			dp[i] = dp[i-1]
		}
		if isValidDecode(s[i-1 : i+1]) {
			if i > 1 {
				dp[i] += dp[i-2]
			} else {
				dp[i] += 1
			}
		}
	}
	return dp[len(s)-1]
}

func isValidDecode(s string) bool {
	i, _ := strconv.ParseInt(s, 10, 32)
	if i > 9 && i < 27 {
		return true
	}
	return false
}

func NumDecoding() {
	out := numDecodings("27")
	fmt.Printf("%v", out)
}
