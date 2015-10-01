package deque

import (
	"container/list"
	"sync"
)

type Deque struct {
	mutex   sync.RWMutex
	data    *list.List
	capcity int
}

func NewDeque() *Deque {
	return NewCappedDeque(-1)
}

func NewCappedDeque(cap int) *Deque {
	return &Deque{
		data:    list.New(),
		capcity: cap,
	}
}

func (d *Deque) Push(item interface{}) bool {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	if d.capcity > 0 && d.data.Len() >= d.capcity {
		return false
	}
	d.data.PushBack(item)
	return true
}

func (d *Deque) Pop() interface{} {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	var item interface{} = nil
	if d.capcity > 0 && d.data.Len() <= 0 {
		return item
	}
	var lastItem *list.Element = d.data.Back()
	if lastItem != nil {
		item = d.data.Remove(lastItem)
	}
	return item
}

func (d *Deque) Len() int {
	return d.data.Len()
}
