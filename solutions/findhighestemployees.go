package solutions

import (
	"fmt"
	"slices"
)

// https://leetcode.com/problems/high-access-employees/
func findHighAccessEmployees(access_times [][]string) []string {
	ret := make(map[string]bool)
	empMap := make(map[string][]int)
	for _, at := range access_times {
		empMap[at[0]] = append(empMap[at[0]], converToMin(at[1]))
	}
	for k, v := range empMap {
		slices.Sort(v)
		for i := 0; i < len(v)-2; i++ {
			if v[i+2]-v[i] < 60 {
				ret[k] = true
			}
		}
	}

	var slice []string
	for s := range ret {
		slice = append(slice, s)
	}

	return slice
}

func converToMin(t string) int {
	return ((int(t[0]-'0')*10)+int(t[1]-'0'))*60 + // hours
		((int(t[2]-'0') * 10) + int(t[3]-'0')) // min
}

func FindHighAccessEmployees() {
	fmt.Print(findHighAccessEmployees([][]string{}))
}
