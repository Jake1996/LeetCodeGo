package solutions

import (
	"container/list"
	"fmt"
)

// https://leetcode.com/problems/largest-rectangle-in-histogram/description/
func largestRectangleArea(heights []int) int {
	n := len(heights)
	minLeft := make([]int, n)
	minRight := make([]int, n)
	for i := range n {
		minLeft[i] = -1
		minRight[i] = n
	}
	stack := list.New()
	for i := 0; i < n; i++ {
		for stack.Len() > 0 {
			top := stack.Back().Value.(int)
			if heights[top] > heights[i] {
				minRight[top] = i
				stack.Remove(stack.Back())
			} else {
				minLeft[i] = top
				break
			}
		}
		stack.PushBack(i)
	}
	maxArea := 0
	for i := range n {
		maxArea = max(maxArea, heights[i]*(minRight[i]-minLeft[i]-1))
	}
	return maxArea
}

func LargestRectangleArea() {
	fmt.Print(largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
}
