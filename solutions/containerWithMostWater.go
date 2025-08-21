package solutions

import "fmt"

// https://leetcode.com/problems/container-with-most-water/description/
func maxArea(height []int) int {
	var i, j, m int
	j = len(height) - 1
	for i < j {
		m = max(m, (j-i)*min(height[j], height[i]))
		if height[i] > height[j] {
			j--
		} else {
			i++
		}
	}
	return m
}

func MaxArea() {
	o := maxArea([]int{1, 3, 1, 3, 5, 3, 6, 7})
	fmt.Printf("%v", o)
}
