package solutions

import (
	"fmt"
	"strconv"
	"strings"
)

// https://leetcode.com/problems/subdomain-visit-count/
func subdomainVisits(cpdomains []string) []string {
	mapCount := make(map[string]int)
	for _, st := range cpdomains {
		initialArr := strings.Split(st, " ")
		num, _ := strconv.ParseInt(initialArr[0], 10, 32)
		numInt := int(num)
		arr := strings.Split(initialArr[1], ".")
		key := ""
		for i := len(arr) - 1; i >= 0; i-- {
			if len(key) > 0 {
				key = arr[i] + "." + key
			} else {
				key = arr[i]
			}
			mapCount[key] = mapCount[key] + numInt
		}
	}
	out := []string{}
	for k := range mapCount {
		out = append(out, fmt.Sprintf("%v %v", mapCount[k], k))
	}
	return out
}

func SubDomainInVisit() {
	fmt.Print(subdomainVisits([]string{}))
}
