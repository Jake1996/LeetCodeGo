package solutions

import "fmt"

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
	var outbi []byte
	for i := range 26 {
		for range bi[i] {
			outbi = append(outbi, byte('a'+rune(i)))
		}
	}
	return string(outbi)
}

func GroupAnagrams() {
	out := groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	fmt.Printf("%v", out)
}
