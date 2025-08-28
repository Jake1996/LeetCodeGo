package solutions

import (
	"fmt"
	"slices"
	"strconv"
)

// https://leetcode.com/problems/high-access-employees/
func findHighAccessEmployees(access_times [][]string) []string {
	ret := make(map[string]bool)
	empMap := make(map[string][]int64)
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

func converToMin(t string) int64 {
	hr, _ := strconv.ParseInt(t[0:2], 10, 32)
	min, _ := strconv.ParseInt(t[2:], 10, 32)

	return hr*60 + min
}

func FindHighAccessEmployees() {
	fmt.Print(findHighAccessEmployees([][]string{}))
}
