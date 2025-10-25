package solutions

import (
	"container/list"
)

// https://leetcode.com/problems/all-oone-data-structure/description/
type Value struct {
	val   string
	count int
}

type AllOne struct {
	// Doubly linked list with max at the front and min at the back
	list *list.List
	hm   map[string]*list.Element
}

func Constructor() AllOne {
	ret := AllOne{
		list: list.New(),
		hm:   make(map[string]*list.Element),
	}
	return ret
}

func (ao *AllOne) Inc(key string) {
	e := ao.hm[key]
	if e == nil {
		v := &Value{
			val:   key,
			count: 1,
		}
		element := ao.list.PushBack(v)
		ao.hm[key] = element
	} else {
		val := e.Value.(*Value)
		val.count = val.count + 1
		for e.Prev() != nil && e.Prev().Value.(*Value).count < val.count {
			ao.list.MoveBefore(e, e.Prev())
		}
	}
}

func (ao *AllOne) Dec(key string) {
	e := ao.hm[key]
	if e == nil {
		return
	} else {
		e := ao.hm[key]
		val := e.Value.(*Value)
		if val.count == 1 {
			ao.list.Remove(e)
			delete(ao.hm, key)
		}
		val.count = val.count - 1
		for e.Next() != nil && e.Next().Value.(*Value).count > val.count {
			ao.list.MoveAfter(e, e.Next())
		}
	}
}

func (ao *AllOne) GetMaxKey() string {
	if ao.list.Front() != nil {
		return ao.list.Front().Value.(*Value).val
	}
	return ""
}

func (ao *AllOne) GetMinKey() string {
	if ao.list.Back() != nil {
		return ao.list.Back().Value.(*Value).val
	}
	return ""
}

/**
 * Your AllOne object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Inc(key);
 * obj.Dec(key);
 * param_3 := obj.GetMaxKey();
 * param_4 := obj.GetMinKey();
 */
