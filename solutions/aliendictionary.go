package solutions

import (
	"container/list"
	"fmt"
)

func alientDictionary(words []string) string {
	adjGraph := make(map[byte][]byte)
	inDegree := make(map[byte]int)

	i := 0
	for i < len(words)-1 {
		if i == len(words) {
			break
		}
		j := 0
		for j < len(words[i]) && j < len(words[i+1]) {
			if words[i][j] != words[i+1][j] {
				adjGraph[words[i][j]] = append(adjGraph[words[i][j]], words[i+1][j])
				if _, ok := inDegree[words[i][j]]; !ok {
					inDegree[words[i][j]] = 0
				}
				inDegree[words[i+1][j]]++
				break
			}
			j++
		}
		i++
	}
	queue := list.New()
	for k, v := range inDegree {
		if v == 0 {
			queue.PushBack(k)
		}
	}
	out := []byte{}
	for queue.Len() > 0 {
		el := queue.Front()
		queue.Remove(el)
		ch := el.Value.(byte)
		out = append(out, ch)
		for _, v := range adjGraph[ch] {
			inDegree[v]--
			if inDegree[v] == 0 {
				queue.PushBack(v)
			}
		}
	}
	return string(out)
}

func AlienDictionary() {
	// fmt.Printf("%v", alientDictionary([]string{"baa", "abcd", "abca", "cab", "cad"}))
	fmt.Printf("%v", alientDictionary([]string{"wrt", "wrf", "er", "ett", "rftt"}))
}
