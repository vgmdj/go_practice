package LRU_Cache

type LRUCache2 struct {
	head     *node
	tail     *node
	capacity int
	selector map[int]*node
}

type node struct {
	previous *node
	next     *node
	key      int
	value    int
}

func Constructor2(Capacity int) LRUCache2 {
	head := &node{}
	tail := &node{}

	head.next = tail
	tail.previous = head

	return LRUCache2{
		head:     head,
		tail:     tail,
		capacity: Capacity,
		selector: make(map[int]*node),
	}
}

func (l *LRUCache2) Get(key int) int {
	n, ok := l.selector[key]
	if !ok {
		return -1
	}

	l.removeNode(n)
	l.Put(key, n.value)

	return n.value
}

func (l *LRUCache2) Put(key, value int) {
	n, ok := l.selector[key]
	if ok {
		l.removeNode(n)
		l.Put(key, value)
		return
	}

	n = &node{
		previous: l.tail.previous,
		next:     l.tail,
		key:      key,
		value:    value,
	}

	l.tail.previous.next = n
	l.tail.previous = n

	l.selector[key] = n

	if len(l.selector) <= l.capacity {
		return
	}

	l.removeNode(l.head.next)

}

func (l *LRUCache2) removeNode(n *node) {
	delete(l.selector, n.key)

	n.previous.next = n.next
	n.next.previous = n.previous

}
