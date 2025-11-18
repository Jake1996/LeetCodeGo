package solutions

import "fmt"

// https://leetcode.com/problems/continuous-subarray-sum/
func checkSubarraySum(nums []int, k int) bool {
	modmap := make(map[int]int, len(nums))
	modmap[nums[0]%k] = 0
	sum := nums[0]
	for i := 1; i < len(nums); i++ {
		sum += nums[i]
		if sum%k == 0 {
			return true
		}
		if v, ok := modmap[sum%k]; ok && i-v > 1 {
			return true
		} else if !ok {
			modmap[sum%k] = i
		}
	}

	return false
}

func CheckSubarraySum() {
	out := checkSubarraySum([]int{23, 2, 4, 6, 7}, 6)
	fmt.Printf("%v", out)
}
