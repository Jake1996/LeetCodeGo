package solutions

import "fmt"

func searchInsert(nums []int, target int) int {
	low := 0
	max := len(nums) - 1

	for low <= max {
		mid := (low + max) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			max = mid - 1
		}
	}
	return low
}

func SearchInsert() {
	fmt.Print(searchInsert([]int{0, 2, 3, 4, 5}, 4))
}
