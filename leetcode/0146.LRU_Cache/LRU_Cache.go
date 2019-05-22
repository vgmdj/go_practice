package LRU_Cache

import (
	"container/list"
)

type LRUCache struct {
	capacity int
	queue    *list.List
	selector map[int]*list.Element
}

type KV struct {
	key   int
	value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		queue:    list.New(),
		selector: make(map[int]*list.Element, 10),
	}
}

func (l *LRUCache) Get(key int) int {
	e, ok := l.selector[key]
	if !ok {
		return -1
	}

	l.queue.MoveToBack(e)

	return e.Value.(KV).value

}

func (l *LRUCache) Put(key int, value int) {
	if l.exist(key, value) {
		return
	}

	l.putNew(key, value)

}

func (l *LRUCache) exist(key int, value int) bool {
	e, ok := l.selector[key]
	if !ok {
		return false
	}

	if e.Value.(KV).value != value {
		l.queue.Remove(e)
		return false
	}

	l.queue.MoveToBack(e)
	return true

}

func (l *LRUCache) putNew(key int, value int) {
	e := KV{
		key:   key,
		value: value,
	}

	if l.queue.Len() >= l.capacity {
		old := l.queue.Front()
		delete(l.selector, old.Value.(KV).key)
		l.queue.Remove(old)

	}

	l.queue.PushBack(e)
	l.selector[key] = l.queue.Back()

}
