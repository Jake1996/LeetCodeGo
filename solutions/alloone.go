package solutions

import (
	"container/list"
)

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

func (this *AllOne) Inc(key string) {
	e := this.hm[key]
	if e == nil {
		v := &Value{
			val:   key,
			count: 1,
		}
		element := this.list.PushBack(v)
		this.hm[key] = element
	} else {
		val := e.Value.(*Value)
		val.count = val.count + 1
		for e.Prev() != nil && e.Prev().Value.(*Value).count < val.count {
			this.list.MoveBefore(e, e.Prev())
		}
	}
}

func (this *AllOne) Dec(key string) {
	e := this.hm[key]
	if e == nil {
		return
	} else {
		e := this.hm[key]
		val := e.Value.(*Value)
		if val.count == 1 {
			this.list.Remove(e)
			delete(this.hm, key)
		}
		val.count = val.count - 1
		for e.Next() != nil && e.Next().Value.(*Value).count > val.count {
			this.list.MoveAfter(e, e.Next())
		}
	}
}

func (this *AllOne) GetMaxKey() string {
	if this.list.Front() != nil {
		return this.list.Front().Value.(*Value).val
	}
	return ""
}

func (this *AllOne) GetMinKey() string {
	if this.list.Back() != nil {
		return this.list.Back().Value.(*Value).val
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
