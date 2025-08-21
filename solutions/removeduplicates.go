package solutions

import "fmt"

func removeDuplicates(nums []int) int {
	var i, j int
	for j < len(nums) {
		if nums[i] == nums[j] {
			j++
		} else {
			i++
			nums[i] = nums[j]
			j++
		}
	}
	return i + 1
}

func RemoveDuplicates() {
	out := removeDuplicates([]int{})
	fmt.Print(out)
}
