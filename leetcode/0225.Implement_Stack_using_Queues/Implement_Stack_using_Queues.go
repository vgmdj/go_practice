package Implement_Stack_using_Queues

import "container/list"

type MyStack struct {
	queue *MyQueue
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{
		queue: NewQueue(),
	}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.queue.Push(x)

	length := this.queue.Size()

	for length > 1 {
		tmp := this.queue.Pop()
		this.queue.Push(tmp)
		length--
	}

}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	return this.queue.Pop()

}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.queue.Peek()
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.queue.IsEmpty()
}

type MyQueue struct {
	l *list.List
}

func NewQueue() *MyQueue {
	return &MyQueue{
		l: list.New(),
	}
}

func (q *MyQueue) Push(x int) {
	q.l.PushBack(x)
}

func (q *MyQueue) Pop() int {
	if q.l.Len() == 0 {
		return -1
	}

	e := q.l.Front()
	q.l.Remove(e)
	return e.Value.(int)
}

func (q *MyQueue) Peek() int {
	if q.l.Len() == 0 {
		return -1
	}

	return q.l.Front().Value.(int)
}

func (q *MyQueue) Size() int {
	return q.l.Len()
}

func (q *MyQueue) IsEmpty() bool {
	return q.l.Len() == 0
}
