package solutions

import (
	"fmt"
	"math"
)

// https://leetcode.com/problems/maximum-product-of-three-numbers/
func maximumProduct(nums []int) int {
	minArr := []int{math.MaxInt, math.MaxInt}              // 0 < 1
	maxArr := []int{math.MinInt, math.MinInt, math.MinInt} // 0 > 1 > 2
	for _, num := range nums {
		if num > maxArr[0] {
			maxArr[2] = maxArr[1]
			maxArr[1] = maxArr[0]
			maxArr[0] = num
		} else if num > maxArr[1] {
			maxArr[2] = maxArr[1]
			maxArr[1] = num
		} else if num > maxArr[2] {
			maxArr[2] = num
		}
		if num < minArr[0] {
			minArr[1] = minArr[0]
			minArr[0] = num
		} else if num < minArr[1] {
			minArr[1] = num
		}
	}
	return max(maxArr[0]*maxArr[1]*maxArr[2], minArr[0]*minArr[1]*maxArr[0])
}

func MaximumProduct() {
	out := maximumProduct([]int{1, 2, 3, 4})
	fmt.Printf("%v", out)
}
