package Design_Front_Middle_Back_Queue

type List struct {
	Val  int
	Pre  *List
	Next *List
}

type FrontMiddleBackQueue struct {
	head   *List
	tail   *List
	mid    *List
	length int
}

func Constructor() FrontMiddleBackQueue {
	head, tail := &List{Val: -1}, &List{Val: -1}
	head.Next = tail
	tail.Pre = head

	return FrontMiddleBackQueue{
		head:   head,
		mid:    head,
		tail:   tail,
		length: 0,
	}
}

func (this *FrontMiddleBackQueue) PushFront(val int) {
	this.length++
	node := &List{
		Val:  val,
		Pre:  this.head,
		Next: this.head.Next,
	}
	this.head.Next.Pre = node
	this.head.Next = node

	if this.length == 1 {
		this.mid = node
		return
	}

	if this.length&1 == 0 {
		this.mid = this.mid.Pre
	}

}

func (this *FrontMiddleBackQueue) PushMiddle(val int) {
	preNode := this.mid
	if this.length&1 == 1 {
		preNode = this.mid.Pre
	}

	this.length++
	node := &List{
		Val:  val,
		Pre:  preNode,
		Next: preNode.Next,
	}
	preNode.Next.Pre = node
	preNode.Next = node

	this.mid = node

}

func (this *FrontMiddleBackQueue) PushBack(val int) {
	this.length++
	node := &List{
		Val:  val,
		Pre:  this.tail.Pre,
		Next: this.tail,
	}
	this.tail.Pre.Next = node
	this.tail.Pre = node

	if this.length == 1 {
		this.mid = node
		return
	}

	if this.length&1 == 1 {
		this.mid = this.mid.Next
	}

}

func (this *FrontMiddleBackQueue) PopFront() int {
	if this.length == 0 {
		return -1
	}

	this.length--
	node := this.head.Next
	this.head.Next = node.Next
	node.Next.Pre = this.head

	if this.length == 0 {
		this.mid = this.head

	} else if this.length&1 == 1 {
		this.mid = this.mid.Next

	}

	return node.Val

}

func (this *FrontMiddleBackQueue) PopMiddle() int {
	if this.length == 0 {
		return -1
	}

	this.length--
	node := this.mid
	this.mid.Pre.Next = node.Next
	this.mid.Next.Pre = node.Pre

	if this.length&1 == 0 {
		this.mid = node.Pre

	} else {
		this.mid = node.Next
	}

	return node.Val

}

func (this *FrontMiddleBackQueue) PopBack() int {
	if this.length == 0 {
		return -1
	}

	this.length--
	node := this.tail.Pre
	this.tail.Pre = node.Pre
	node.Pre.Next = this.tail

	if this.length&1 == 0 {
		this.mid = this.mid.Pre

	}

	return node.Val

}

/**
 * Your FrontMiddleBackQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PushFront(val);
 * obj.PushMiddle(val);
 * obj.PushBack(val);
 * param_4 := obj.PopFront();
 * param_5 := obj.PopMiddle();
 * param_6 := obj.PopBack();
 */
