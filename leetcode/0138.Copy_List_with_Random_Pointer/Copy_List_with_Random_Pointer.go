package Copy_List_with_Random_Pointer

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	refMap := make(map[*Node]*Node)
	newHead := &Node{
		Val: head.Val,
	}
	refMap[head] = newHead

	preNew := newHead
	cNew, cOld := newHead.Next, head.Next
	for cOld != nil {
		cNew = &Node{
			Val: cOld.Val,
		}
		refMap[cOld] = cNew
		cOld = cOld.Next

		preNew.Next = cNew
		preNew = preNew.Next
	}

	cNew, cOld = newHead, head
	for cOld != nil {
		cNew.Random = refMap[cOld.Random]
		cOld = cOld.Next
		cNew = cNew.Next
	}

	return newHead

}
