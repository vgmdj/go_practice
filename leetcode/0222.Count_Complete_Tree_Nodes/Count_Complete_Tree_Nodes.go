package Count_Complete_Tree_Nodes

import "sort"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// 根据最左叶子，算树的高度
	level := 0
	for node := root; node.Left != nil; node = node.Left {
		level++
	}

	// 在最后一层中，利用二分搜索查找第一个空节点
	// k的二进制位，代表从根节点到k节点的路径，
	// 最高位代表是否是在第h层，其余各位从高到低，0代表要移动到左子树，1代表要移动到右子树
	return sort.Search(1<<(level+1), func(k int) bool {
		if k <= 1<<level {
			return false
		}
		bits := 1 << (level - 1)
		node := root
		for node != nil && bits > 0 {
			if bits&k == 0 {
				node = node.Left
			} else {
				node = node.Right
			}
			bits >>= 1
		}
		return node == nil
	}) - 1
}
