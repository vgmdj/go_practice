package binary_search_tree

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

func (root *AVLTree) searchNodeAndParent(val int) (c *AVLTree, p *AVLTree) {
	c = root

	for c != nil {
		if c.data == val {
			return c, p
		}
		p = c

		if c.data < val {
			c = c.right

		} else {
			c = c.left
		}

	}

	return c, p
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
	c, p := root.searchNodeAndParent(val)
	if c == nil {
		return root
	}

	successor := new(AVLTree)
	if c.left == nil {
		successor = c.right

	} else if c.right == nil {
		successor = c.left

	} else {
		// find the left's max child node
		lp, lmax := c, c.left
		for lmax.right != nil {
			lp = lmax
			lmax = lmax.right
		}

		// delete the left's max child node
		if lp.left == lmax {
			lp.left = lmax.left

		} else {
			lp.right = lmax.right

		}

		successor = lmax
		successor.left = c.left
		successor.right = c.right

	}

	if p == nil {
		return successor
	}

	if p.left == c {
		p.left = successor

	} else {
		p.right = successor

	}

	return root
}

func (root *AVLTree) resetHeight(val int) {
	if root == nil || root.data == val {
		return
	}

	root.height = max(root.left.getHeight(), root.right.getHeight()) + 1

	if val > root.data {
		root.right.resetHeight(val)
		return
	}

	root.left.resetHeight(val)

}

func (root *AVLTree) keepBalance() *AVLTree {
	if root == nil {
		return root
	}

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
