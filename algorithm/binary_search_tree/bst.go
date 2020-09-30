package binary_search_tree

import "math"

type BinarySearchTree struct {
	data  int
	left  *BinarySearchTree
	right *BinarySearchTree
}

func (root *BinarySearchTree) ToArray() []int {
	result := make([]int, 0)

	if root == nil {
		return result
	}

	dfsHelper(&result, root, 0)

	return result
}

func dfsHelper(result *[]int, root *BinarySearchTree, index int) {
	if root == nil {
		return
	}

	if len(*result) < index+1 {
		*result = append(*result, make([]int, index+1-len(*result))...)
	}

	(*result)[index] = root.data

	dfsHelper(result, root.left, index*2+1)
	dfsHelper(result, root.right, index*2+2)

}

func NewBinarySearchTree(val int) *BinarySearchTree {
	return &BinarySearchTree{
		data: val,
	}
}

func (root *BinarySearchTree) IsBst() bool {
	return isBstHelper(root, math.MinInt32, math.MaxInt32)
}

func isBstHelper(root *BinarySearchTree, lower, upper int) bool {
	if root == nil {
		return true
	}

	if root.data <= lower || root.data >= upper {
		return false
	}

	return isBstHelper(root.left, lower, root.data) && isBstHelper(root.right, root.data, upper)

}

func (root *BinarySearchTree) AddNode(val int) {
	if root.data == val {
		return
	}

	if root.data < val {
		if root.right == nil {
			root.right = &BinarySearchTree{
				data: val,
			}

		} else {
			root.right.AddNode(val)

		}

	} else {
		if root.left == nil {
			root.left = &BinarySearchTree{
				data: val,
			}

		} else {
			root.left.AddNode(val)

		}
	}

}

// RemoveNode 先找到要删除的node节点
// 如果节点左为空，则将节点右上提
// 如果节点右为空，则将节点左上提
// 如果左右都不为空，则找到比节点大的最小值，将最小值上提（需注意最小值的右节点处理）
// 或者是找到比节点小的最大值，将这个最大值上提（需注意最大值的左节点处理）
func (root *BinarySearchTree) RemoveNode(val int) {
	res := root.removeNodeHelper(val)
	root.data = res.data
	root.left = res.left
	root.right = res.right
}

func (root *BinarySearchTree) removeNodeHelper(val int) *BinarySearchTree {
	if root == nil {
		return nil
	}

	if root.data < val {
		root.right = root.right.removeNodeHelper(val)
		return root
	}

	if root.data > val {
		root.left = root.left.removeNodeHelper(val)
		return root
	}

	// root.data == val
	if root.left == nil {
		return root.right

	} else if root.right == nil {
		return root.left

	} else {
		successor := root.right.MinNode()
		successor.left = root.left
		return root.right
	}

}

func (root *BinarySearchTree) Search(val int) *BinarySearchTree {
	if root == nil {
		return nil
	}

	if val == root.data {
		return root
	}

	if val > root.data {
		return root.right.Search(val)
	}

	return root.left.Search(val)
}

func (root *BinarySearchTree) MinNode() *BinarySearchTree {
	if root == nil {
		return nil
	}

	if root.left == nil {
		return root
	}

	return root.left.MinNode()
}

func (root *BinarySearchTree) MaxNode() *BinarySearchTree {
	if root == nil {
		return nil
	}

	if root.right == nil {
		return root
	}

	return root.right.MaxNode()
}
