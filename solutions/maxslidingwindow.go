package solutions

import (
	"container/heap"
	"fmt"
)

type MaxHeap []int

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i] > h[j] // This makes it a MAX HEAP
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func maxSlidingWindow(nums []int, k int) []int {
	h := &MaxHeap{}
	heap.Init(h) // Initialize the heap structure
	// var left, right int

	// for right < len(nums) {

	// }

	for _, i := range nums {
		heap.Push(h, i)
	}

	return []int{}
}

func MaxSlidingWindow() {
	o := maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3)
	fmt.Printf("%v", o)
}
