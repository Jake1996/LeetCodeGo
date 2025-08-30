package solutions

import "fmt"

// https://leetcode.com/problems/longest-increasing-subsequence/description/
func lengthOfLIS(nums []int) int {
	inc := make([]int, len(nums))
	size := 0
	for _, num := range nums {
		i := 0
		j := size
		for i != j {
			mid := (i + j) / 2
			if num > inc[mid] {
				i = mid + 1
			} else {
				j = mid // Do not go beyond size
			}
		}
		inc[i] = num
		if i == size {
			size++
		}
	}
	return size
}

func LongestIncreasingSequence() {
	fmt.Print(lengthOfLIS([]int{0, 1, 0, 3, 2, 3}))
}
