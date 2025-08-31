package solutions

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

// https://leetcode.com/problems/alert-using-same-key-card-three-or-more-times-in-a-one-hour-period/
func alertNames(keyName []string, keyTime []string) []string {
	out := []string{}
	timeMap := map[string][]int{}
	for i := range len(keyName) {
		timeMap[keyName[i]] = append(timeMap[keyName[i]], timeToInt(keyTime[i]))
	}
	for k, v := range timeMap {
		if threeOrMore(v) {
			out = append(out, k)
		}
	}
	slices.Sort(out)
	return out
}

func threeOrMore(arr []int) bool {
	slices.Sort(arr)
	for i := 0; i < len(arr)-2; i++ {
		if arr[i+2]-arr[i] <= 60 {
			return true
		}
	}
	return false
}

func timeToInt(s string) int {
	timeArray := strings.Split(s, ":")
	h, _ := strconv.ParseInt(timeArray[0], 10, 64)
	m, _ := strconv.ParseInt(timeArray[1], 10, 64)
	return int(h*60 + m)
}

func AlertNames() {
	fmt.Print(alertNames([]string{}, []string{}))
}
