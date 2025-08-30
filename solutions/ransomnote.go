package solutions

import "fmt"

// https://leetcode.com/problems/ransom-note/
func canConstruct(ransomNote string, magazine string) bool {
	letters := make([]int, 26)
	for _, c := range magazine {
		letters[c-'a']++
	}

	for _, c := range ransomNote {
		if letters[c-'a'] <= 0 {
			return false
		}
		letters[c-'a']--
	}

	return true
}

func CanConstruct() {
	fmt.Print(canConstruct("aa", "aa"))
}
