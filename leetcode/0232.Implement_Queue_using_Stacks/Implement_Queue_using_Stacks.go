package Implement_Queue_using_Stacks

import "container/list"

type MyQueue struct {
	front int
	s1    *MyStack
	s2    *MyStack
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		s1: NewStack(),
		s2: NewStack(),
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	if this.s1.IsEmpty() {
		this.front = x
	}
	this.s1.Push(x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if this.s2.IsEmpty() {
		for !this.s1.IsEmpty() {
			this.s2.Push(this.s1.Pop())
		}
	}

	return this.s2.Pop()

}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if !this.s2.IsEmpty() {
		return this.s2.Peek()
	}

	return this.front
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.s1.IsEmpty() && this.s2.IsEmpty()
}

type MyStack struct {
	l *list.List
}

func NewStack() *MyStack {
	return &MyStack{
		l: list.New(),
	}
}

func (s *MyStack) Push(x int) {
	s.l.PushBack(x)
}

func (s *MyStack) Pop() int {
	if s.l.Len() == 0 {
		return -1
	}

	e := s.l.Back()
	s.l.Remove(e)
	return e.Value.(int)
}

func (s *MyStack) Peek() int {
	if s.l.Len() == 0 {
		return -1
	}

	return s.l.Back().Value.(int)
}

func (s *MyStack) Size() int {
	return s.l.Len()
}

func (s *MyStack) IsEmpty() bool {
	return s.l.Len() == 0
}
