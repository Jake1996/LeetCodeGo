package solutions

import (
	"container/heap"
	"fmt"
)

type Item struct {
	Value      int
	Additional any
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Value > pq[j].Value
}

func (pq *PriorityQueue) Push(e any) {
	*pq = append(*pq, e.(*Item))
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	last := old[n-1]
	*pq = old[0 : n-1]
	return last
}

func TestPQ() {
	p := &PriorityQueue{}
	heap.Init(p)
	heap.Push(p, &Item{Value: 1})
	heap.Push(p, &Item{Value: 4})
	heap.Push(p, &Item{Value: 3})
	heap.Push(p, &Item{Value: 12})
	heap.Push(p, &Item{Value: 6})
	heap.Push(p, &Item{Value: 7})
	fmt.Print(heap.Pop(p))
	fmt.Print(heap.Pop(p))
	fmt.Print(heap.Pop(p))
	fmt.Print(heap.Pop(p))
}
