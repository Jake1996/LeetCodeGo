package solutions

import (
	"fmt"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	strMap := make(map[string][]string)

	for _, str := range strs {
		sortedStr := sortedStr(str)
		strMap[sortedStr] = append(strMap[sortedStr], str)
	}
	var out [][]string
	for k := range strMap {
		out = append(out, strMap[k])
	}
	return out
}

func sortedStr(s string) string {
	bi := make([]int, 26)
	for _, c := range s {
		bi[c-'a']++
	}
	sb := strings.Builder{}
	for i := range 26 {
		for range bi[i] {
			sb.WriteByte(byte('a' + i))
		}
	}
	return sb.String()
}

func GroupAnagrams() {
	out := groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	fmt.Printf("%v", out)
}
