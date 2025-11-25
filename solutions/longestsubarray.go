package solutions

import "container/list"

// https://leetcode.com/problems/longest-continuous-subarray-with-absolute-diff-less-than-or-equal-to-limit/
func LongestSubarray(nums []int, limit int) int {

	minQ := make([]int, 0, len(nums))
	maxQ := make([]int, 0, len(nums))
	i := 0
	res := 0
	for j := 0; j < len(nums); j++ {
		for len(minQ) > 0 && minQ[len(minQ)-1] > nums[j] {
			minQ = minQ[:len(minQ)-1]
		}
		minQ = append(minQ, nums[j])

		for len(maxQ) > 0 && maxQ[len(maxQ)-1] < nums[j] {
			maxQ = maxQ[:len(maxQ)-1]
		}
		maxQ = append(maxQ, nums[j])

		for maxQ[0]-minQ[0] > limit {
			if nums[i] == minQ[0] {
				minQ = minQ[1:]
			}
			if nums[i] == maxQ[0] {
				maxQ = maxQ[1:]
			}
			i++
		}
		res = max(res, j-i+1)
	}

	return res
}

func LongestSubarray2(nums []int, limit int) int {
	res := 1

	start := 0
	minList := list.New()
	maxList := list.New()
	for cur, num := range nums {
		for maxList.Len() > 0 && maxList.Back().Value.(int) < num {
			maxList.Remove(maxList.Back())
		}
		maxList.PushBack(num)
		for minList.Len() > 0 && minList.Back().Value.(int) > num {
			minList.Remove(minList.Back())
		}
		minList.PushBack(num)

		for maxList.Front().Value.(int)-minList.Front().Value.(int) > limit {
			if maxList.Front().Value.(int) == nums[start] {
				maxList.Remove(maxList.Front())
			}
			if minList.Front().Value.(int) == nums[start] {
				minList.Remove(minList.Front())
			}
			start++
		}
		res = max(res, cur-start+1)
	}
	return res
}
