package solutions

import (
	"fmt"
	"slices"
	"sort"
)

// https://leetcode.com/problems/custom-sort-string/description/
func customSortString(order string, s string) string {
	orderMap := make(map[rune]int)

	for i, c := range order {
		orderMap[c] = i
	}
	st := []rune(s)
	slices.SortFunc(st, func(a, b rune) int {
		val1 := orderMap[a]
		val2 := orderMap[b]
		return val1 - val2
	})
	return string(st)
}

func alternative(order string, s string) string {
	m := map[byte]int{}
	for i := range order {
		m[order[i]] = i
	}
	bs := []byte(s)
	sort.Slice(bs, func(i, j int) bool {
		return m[bs[i]] < m[bs[j]]
	})
	return string(bs)
}

func CustomSortString() {
	out := customSortString("cba", "abcd")
	out2 := alternative("cba", "abcd")
	fmt.Printf("%v %v", out, out2)
}
