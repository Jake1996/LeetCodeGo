package solutions

// https://leetcode.com/problems/squares-of-a-sorted-array/
func SortedSquares(nums []int) []int {
	var i, j, k int
	out := make([]int, len(nums))
	for i = 0; i < len(nums); i++ {
		if nums[i] >= 0 {
			break
		}
	}
	j = i - 1
	for i < len(nums) && j >= 0 {
		is := nums[i] * nums[i]
		js := nums[j] * nums[j]
		if is < js {
			out[k] = is
			i++
			k++
		} else {
			out[k] = js
			j--
			k++
		}
	}
	for j >= 0 {
		out[k] = nums[j] * nums[j]
		j--
		k++
	}
	for i < len(nums) {
		out[k] = nums[i] * nums[i]
		i++
		k++
	}
	return out
}
