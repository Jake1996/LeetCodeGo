package solutions

import (
	"fmt"
	"slices"
	"strings"
)

func reverseWords(s string) string {
	s = strings.Trim(s, " ")
	arrs := strings.Split(s, " ")
	newArrs := []string{}
	for _, arr := range arrs {
		if len(strings.Trim(arr, " ")) != 0 {
			newArrs = append(newArrs, strings.Trim(arr, " "))
		}
	}
	slices.Reverse(newArrs)
	return strings.Join(newArrs, " ")
}

func ReverseWords() {
	fmt.Print(reverseWords("abs ic bcd"))
}
