package solutions

import (
	"fmt"
)

// https://leetcode.com/problems/bulls-and-cows/description/

func getHint(secret string, guess string) string {
	var b, c int
	var sec [10]int
	var gue [10]int
	for i := 0; i < len(secret); i++ {
		if secret[i] == guess[i] {
			b++
		} else {
			sec[secret[i]-'0']++
			gue[guess[i]-'0']++
		}
	}
	for i := range 10 {
		c += min(sec[i], gue[i])
	}

	return fmt.Sprintf("%vA%vB", b, c)
}

func BullsAndCows() {
	o := getHint("1807", "7810")
	fmt.Printf("%v", o)
}
