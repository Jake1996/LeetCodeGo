package solutions

// https://leetcode.com/problems/set-mismatch/
func FindErrorNums(nums []int) []int {
	exists := make(map[int]bool)
	var total int64
	out := []int{0, 0}
	for _, num := range nums {
		if exists[num] {
			out[0] = num
		}
		exists[num] = true
		total += int64(num)
	}
	n := len(nums)
	expected := (n * (n + 1)) / 2
	total -= int64(out[0])
	out[1] = int(int64(expected) - total)
	return out
}
