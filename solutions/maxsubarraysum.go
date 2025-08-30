package solutions

import (
	"fmt"
)

// https://leetcode.com/problems/maximum-subarray/
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	maxSum := nums[0]
	curSum := 0
	for _, r := range nums {
		if curSum < 0 {
			curSum = 0
		}
		curSum += r
		maxSum = max(maxSum, curSum)
	}
	return maxSum
}

func MaxSubArraySum() {
	fmt.Print(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
