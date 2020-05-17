package Design_Linked_List

type NodeList struct {
	Val  int
	Prev *NodeList
	Next *NodeList
}

type MyLinkedList struct {
	head   *NodeList
	tail   *NodeList
	length int
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{
		head:   nil,
		tail:   nil,
		length: 0,
	}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.length {
		return -1
	}

	return this.getNode(index).Val

}

func (this *MyLinkedList) getNode(index int) *NodeList {
	if index < 0 || index >= this.length {
		return nil
	}

	mid := this.length / 2
	if index <= mid {
		return this.getFromHead(index)
	}

	return this.getFromTail(this.length - index - 1)
}

func (this *MyLinkedList) getFromHead(index int) *NodeList {
	res := this.head
	for ; index > 0; index-- {
		res = res.Next
	}
	return res
}

func (this *MyLinkedList) getFromTail(index int) *NodeList {
	res := this.tail
	for ; index > 0; index-- {
		res = res.Prev
	}
	return res
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	newHead := &NodeList{
		Val:  val,
		Prev: nil,
		Next: this.head,
	}

	if this.head != nil {
		this.head.Prev = newHead

	} else {
		this.tail = newHead

	}

	this.head = newHead

	this.length = this.length + 1
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	newTail := &NodeList{
		Val:  val,
		Prev: this.tail,
		Next: nil,
	}

	if this.tail != nil {
		this.tail.Next = newTail

	} else {
		this.head = newTail

	}

	this.tail = newTail

	this.length = this.length + 1

}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.length {
		return
	}

	if index == this.length {
		this.AddAtTail(val)
		return
	}

	if index <= 0 {
		this.AddAtHead(val)
		return
	}

	node := this.getNode(index)
	prev := node.Prev
	newNode := &NodeList{
		Val:  val,
		Prev: prev,
		Next: node,
	}
	prev.Next = newNode
	node.Prev = newNode

	this.length = this.length + 1
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.length {
		return
	}

	node := this.getNode(index)
	if node.Prev == nil {
		this.head = node.Next

	} else {
		node.Prev.Next = node.Next

	}

	if node.Next == nil {
		this.tail = node.Prev

	} else {
		node.Next.Prev = node.Prev

	}

	this.length = this.length - 1

}
