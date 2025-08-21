package solutions

import (
	"fmt"
)

// https://leetcode.com/problems/max-consecutive-ones-iii/

func longestOnes(nums []int, k int) int {
	var m, n, maxVal, t int
	for n < len(nums) {
		if nums[n] == 1 {
			n++
		} else {
			if t < k {
				t++
				n++
			} else {
				m++
				if nums[m-1] == 0 {
					t--
				}
			}
		}
		maxVal = max(maxVal, (n - m))
	}
	return maxVal
}

func Maxconsecutiveones() {
	o := longestOnes([]int{1, 1, 1, 1, 1}, 3)
	fmt.Printf("%v", o)
}
