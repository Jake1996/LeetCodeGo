package solutions

import "fmt"

// https://leetcode.com/problems/subarrays-with-k-different-integers/

func subarraysWithKDistinct(nums []int, k int) int {
	if k == 0 {
		return 1
	}
	var left, right, result, distinct, count int
	var unique map[int]int = make(map[int]int)
	for right < len(nums) {
		if unique[nums[right]] == 0 {
			distinct++
		}
		unique[nums[right]] = unique[nums[right]] + 1
		right++
		for distinct > k {
			count = 0
			if unique[nums[left]] == 1 {
				distinct--
			}
			unique[nums[left]] = unique[nums[left]] - 1
			left++
		}
		if distinct == k {
			for unique[nums[left]] > 1 {
				count++
				unique[nums[left]] = unique[nums[left]] - 1
				left++
			}
			result = result + count + 1
		}
	}
	return result
}

func SubArraysWithKDisctinct() {
	o := subarraysWithKDistinct([]int{1, 2, 1, 2, 3}, 2)
	fmt.Printf("%v", o)
}
