package solutions

import "fmt"

// https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/
func searchRange(nums []int, target int) []int {
	left := findMidMin(nums, target)

	if left == -1 {
		return []int{-1, -1}
	}
	right := findMidMax(nums, target, left)
	return []int{left, right}
}

func findMidMin(nums []int, target int) int {
	var min, max, mid, found int
	max = len(nums) - 1
	found = -1
	for min <= max {
		mid = (max + min) / 2
		if nums[mid] == target {
			found = mid
			max = mid - 1
		} else if nums[mid] > target {
			max = mid - 1
		} else {
			min = mid + 1
		}
	}
	return found
}

func findMidMax(nums []int, target int, start int) int {
	var min, max, mid, found int
	max = len(nums) - 1
	min = start
	found = -1
	for min <= max {
		mid = (max + min) / 2
		if nums[mid] == target {
			found = mid
			min = mid + 1
		} else if nums[mid] > target {
			max = mid - 1
		} else {
			min = mid + 1
		}
	}
	return found
}

func FirstAndLast() {
	fmt.Print(searchRange([]int{}, 0))
}
