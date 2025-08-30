package solutions

import (
	"container/heap"
	"fmt"
)

type ItemTop struct {
	val   int
	freq  int
	index int
}

type PriorityQueueTop []*ItemTop

func (pq PriorityQueueTop) Len() int {
	return len(pq)
}

func (pq PriorityQueueTop) Swap(i, j int) {
	pq[i].index = j
	pq[j].index = i
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq PriorityQueueTop) Less(i, j int) bool {
	return pq[i].freq > pq[j].freq
}

func (pq *PriorityQueueTop) Push(item any) {
	item.(*ItemTop).index = len(*pq)
	*pq = append(*pq, item.(*ItemTop))
}

func (pq *PriorityQueueTop) Pop() any {
	old := *pq
	last := old[len(old)-1]
	*pq = old[:len(old)-1]
	return last
}

func topKFrequent(nums []int, k int) []int {
	queue := make(PriorityQueueTop, 0)
	heap.Init(&queue)
	store := make(map[int]*ItemTop)
	for _, el := range nums {
		if store[el] == nil {
			itel := &ItemTop{
				val:  el,
				freq: 1,
			}
			store[el] = itel
			heap.Push(&queue, itel)
		} else {
			itel := store[el]
			itel.freq = itel.freq + 1
			heap.Fix(&queue, itel.index)
		}
	}
	out := make([]int, 0)
	for range k {
		if queue.Len() > 0 {
			it := heap.Pop(&queue).(*ItemTop)
			out = append(out, it.val)
		}
	}
	return out
}

func TopKFrequent() {
	fmt.Print(topKFrequent([]int{1, 2}, 2))
	fmt.Print(topKFrequentAlternate([]int{1, 2}, 2))
}

func topKFrequentAlternate(nums []int, k int) []int {
	lengthOfNums := len(nums)
	if lengthOfNums == 1 || lengthOfNums == 0 {
		return nums
	}

	counterMap := make(map[int]int)
	for _, num := range nums {
		counterMap[num]++
	}

	bucket := make([][]int, lengthOfNums)
	result := []int{}

	for index, uniqueCounter := range counterMap {
		bucket[uniqueCounter-1] = append(bucket[uniqueCounter-1], index)
	}
	for i := lengthOfNums - 1; i >= 0; i-- {
		if len(bucket[i]) > 0 {
			result = append(result, bucket[i]...)

		}
		if len(result) == k {
			break
		}
	}

	return result
}
