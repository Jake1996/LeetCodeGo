package solutions

import (
	"container/list"
	"fmt"
)

func maxResult(nums []int, k int) int {
	dequeue := list.New()

	longest := make([]int, len(nums))

	longest[0] = nums[0]
	dequeue.PushBack(0)

	for i := 1; i < len(nums); i++ {
		mx := dequeue.Front().Value.(int)
		longest[i] = nums[i] + longest[mx]
		for dequeue.Len() > 0 && longest[dequeue.Back().Value.(int)] <= longest[i] {
			dequeue.Remove(dequeue.Back())
		}
		dequeue.PushBack(i)
		for dequeue.Front().Value.(int) <= i-k {
			dequeue.Remove(dequeue.Front())
		}
	}

	return longest[len(nums)-1]
}

func MaxResult() {
	out := maxResult([]int{1, -1, -2, 4, -7, 3}, 2)
	fmt.Print(out)
}
