package LRU_Cache

type LRUCache2 struct {
	Head     *Node
	Tail     *Node
	Capacity int
	selector map[int]*Node
}

func Constructor2(capacity int) LRUCache2 {
	head, tail := new(Node), new(Node)
	head.Next = tail
	tail.Prev = head

	return LRUCache2{
		Head:     head,
		Tail:     tail,
		Capacity: capacity,
		selector: make(map[int]*Node, 0),
	}

}

func (this *LRUCache2) Get(key int) int {
	n, ok := this.selector[key]
	if !ok {
		return -1
	}

	this.removeNode(n)
	this.Put(key, n.Value)

	return n.Value

}

func (this *LRUCache2) Put(key int, value int) {
	n, ok := this.selector[key]
	if ok {
		this.removeNode(n)
	}

	if len(this.selector) >= this.Capacity {
		this.removeNode(this.Head.Next)
	}

	node := &Node{
		Prev:  this.Tail.Prev,
		Next:  this.Tail,
		Key:   key,
		Value: value,
	}

	node.Prev.Next = node
	node.Next.Prev = node

	this.selector[key] = node

}

func (this *LRUCache2) removeNode(node *Node) {
	if node == nil {
		return
	}

	delete(this.selector, node.Key)

	prev := node.Prev
	next := node.Next
	prev.Next = next
	next.Prev = prev

}

type Node struct {
	Prev  *Node
	Next  *Node
	Key   int
	Value int
}
