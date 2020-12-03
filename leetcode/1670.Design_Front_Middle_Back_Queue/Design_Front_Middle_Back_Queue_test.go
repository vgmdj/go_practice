package Design_Front_Middle_Back_Queue

import "testing"

func TestFrontMiddleBackQueue(t *testing.T) {
	q1 := Constructor()
	// 2 3 1
	q1.PushMiddle(1)
	q1.PushMiddle(2)
	q1.PushMiddle(3)

	// 3 2 1
	t.Log(q1.PopMiddle())
	t.Log(q1.PopMiddle())
	t.Log(q1.PopMiddle())

	t.Log("case 2")


	q := Constructor()

	//5 1 4 3 2
	q.PushFront(1)
	q.PushBack(2)
	q.PushMiddle(3)
	q.PushMiddle(4)
	q.PushFront(5)

	// 5 4 3 2 1
	t.Log(q.PopFront())
	t.Log(q.PopMiddle())
	t.Log(q.PopMiddle())
	t.Log(q.PopBack())
	t.Log(q.PopFront())
}
