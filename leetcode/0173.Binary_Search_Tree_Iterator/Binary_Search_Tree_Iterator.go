package Binary_Search_Tree_Iterator

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	stack *list.List
	pos   *TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	return BSTIterator{
		stack: list.New(),
		pos:   root,
	}
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	for this.pos != nil {
		this.stack.PushBack(this.pos)
		this.pos = this.pos.Left
	}

	node := new(TreeNode)
	if this.stack.Len() != 0 {
		el := this.stack.Back()
		this.stack.Remove(el)
		node = el.Value.(*TreeNode)

		this.pos = node.Right

	}

	return node.Val

}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	if this.stack.Len() != 0 || this.pos != nil {
		return true
	}

	return false
}
