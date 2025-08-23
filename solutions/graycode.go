package solutions

import "fmt"

// https://leetcode.com/problems/gray-code/description/
func grayCode(n int) []int {

	if n == 1 {
		list := []int{}
		list = append(list, 0)
		list = append(list, 1)
		return list
	}

	prev := grayCode(n - 1)
	ans := []int{}
	for i := 0; i < len(prev); i++ {
		ans = append(ans, prev[i]<<1+0)
	}
	for i := len(prev) - 1; i >= 0; i-- {
		ans = append(ans, prev[i]<<1+1)
	}
	return ans
}

func GrayCode() {
	ans := grayCode(5)
	fmt.Printf("%v", ans)
}
