package solutions

import (
	"container/list"
)

type Queue struct {
	data *list.List
}

type Point struct {
	Point int
	Cost  int
}

type Node struct {
	Point      int
	Hops       int
	TotalPrice int
}

func NewQueue() *Queue {
	q := &Queue{
		data: list.New(),
	}
	return q
}

func (q *Queue) IsEmpty() bool {
	return q.data.Len() == 0
}

func (q *Queue) Enqueue(a *Node) {
	q.data.PushBack(a)
}

func (q *Queue) Dequeue() (*Node, bool) {
	if q.data.Len() == 0 {
		return nil, false
	}
	e := q.data.Front()
	q.data.Remove(e)
	return e.Value.(*Node), true
}
