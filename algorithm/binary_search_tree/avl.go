package binary_search_tree

import "math"

type AVLTree struct {
	data   int
	left   *AVLTree
	right  *AVLTree
	height int
}

func (root *AVLTree) getHeight() int {
	if root == nil {
		return 0
	}

	return root.height
}

func (root *AVLTree) addNode(val int) *AVLTree {
	if root == nil {
		return &AVLTree{
			data:   val,
			height: 1,
		}
	}

	if root.data < val {
		root.right = root.right.addNode(val)

	} else if root.data > val {
		root.left = root.left.addNode(val)

	}

	root.height = max(root.left.getHeight(), root.right.getHeight()) + 1
	return root

}

func (root *AVLTree) removeNode(val int) *AVLTree {
	c := root
	p := &AVLTree{
		data: math.MaxInt32,
		left: c,
	}

	for c != nil {
		if c.data == val {
			break
		}
		p = c

		if c.data < val {
			c = c.left

		} else {
			c = c.right
		}

	}

	if c == nil {
		return root
	}

	successor := new(AVLTree)
	if c.left == nil {
		successor = c.right

	} else if c.right == nil {
		successor = c.left

	} else {
		lp, lmax := c, c.left
		for lmax.right != nil {
			lp = lmax
			lmax = lmax.right
		}

		if lp.left == lmax {
			lp.left = nil

		} else {
			lp.right = nil

		}
		successor = lmax




	}

	return root
}

func (root *AVLTree) keepBalance() *AVLTree {

	left, right := root.left, root.right
	if left.getHeight()-right.getHeight() >= 2 {
		if left.left.getHeight() > left.right.getHeight() {
			return root.rightRotate()

		} else {
			return root.leftRightRotate()

		}

	}

	if right.getHeight()-left.getHeight() >= 2 {
		if right.right.getHeight() > right.left.getHeight() {
			return root.leftRotate()

		} else {
			return root.rightLeftRotate()

		}

	}

	return root

}

func (root *AVLTree) leftRotate() *AVLTree {
	right := root.right
	root.right = right.left
	right.left = root

	root.height = max(root.left.getHeight(), root.right.getHeight()) + 1
	right.height = max(right.left.getHeight(), right.right.getHeight()) + 1

	return right

}

func (root *AVLTree) rightRotate() *AVLTree {
	left := root.left
	root.left = left.right
	left.right = root

	root.height = max(root.left.getHeight(), root.right.getHeight()) + 1
	left.height = max(left.left.getHeight(), left.right.getHeight()) + 1

	return left
}

func (root *AVLTree) leftRightRotate() *AVLTree {
	root.left = root.left.leftRotate()
	return root.rightRotate()
}

func (root *AVLTree) rightLeftRotate() *AVLTree {
	root.right = root.right.rightRotate()
	return root.leftRotate()
}

type AVLHelper struct {
	root *AVLTree
}

func NewAVLHelper() *AVLHelper {
	return &AVLHelper{}
}

func (h *AVLHelper) GetAVLTree() *AVLTree {
	return h.root

}

func (h *AVLHelper) AddNode(val int) {
	h.root = h.root.addNode(val)
	h.root = h.root.keepBalance()
}

func (h *AVLHelper) RemoveNode(val int) {
	h.root = h.root.removeNode(val)
	h.root = h.root.keepBalance()
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
