package binary_search_tree

import (
	"errors"
)


var (
	Wrong = errors.New("Wrong")
)

type BiTreeNode struct {
	data  interface{}
	left  *BiTreeNode
	right *BiTreeNode
}

type BiTree struct {
	size    int
	compare func(key1, key2 interface{}) int
	root    *BiTreeNode
}

func (bs *BiTree) Init(args ...interface{}) {
	bs.size = 0
	if len(args) == 1 {
		bs.compare = args[0].(func(key1, key2 interface{}) int)
	}
	return
}

//将新节点插入到二叉树中且node节点的左子结点，新节点的数据为data。若node为nil且二叉树为空树，则新结点作为根结点
func (bs *BiTree) Ins_left(node *BiTreeNode,
	data interface{}) (e error) {

	var pos = &BiTreeNode{data: data}
	if node == nil {
		if bs.size > 0 {
			return Wrong
		}

		bs.root = pos
	} else {
		if node.left != nil {
			return Wrong
		}
		node.left = pos
	}

	bs.size++
	return
}

func (bs *BiTree) Ins_right(node *BiTreeNode, data interface{}) (e error) {
	var pos = &BiTreeNode{data: data}
	if node == nil {
		if bs.size > 0 {
			return Wrong
		}

		bs.root = pos
	} else {
		if node.right != nil {
			return Wrong
		}
		node.right = pos
	}

	bs.size++
	return nil
}

//将指定结点node的左子树移除,node 为nil，以根结点为起始结点
func (bs *BiTree) Rem_left(node *BiTreeNode) {
	if bs.size == 0 {
		return
	}

	var pos **BiTreeNode

	if node == nil {
		pos = &bs.root
	} else {
		pos = &node.left
	}

	//后序遍历删除所有的结点
	if *pos != nil {
		bs.Rem_left(*pos)
		bs.Rem_right(*pos)
		*pos = nil
		bs.size --
	}
	return
}

func (bs *BiTree) Rem_right(node *BiTreeNode) {
	if bs.size == 0 {
		return
	}

	var pos **BiTreeNode

	if node == nil {
		pos = &bs.root
	} else {
		pos = &node.right
	}

	//后序遍历删除所有的结点
	if *pos != nil {
		bs.Rem_left(*pos)
		bs.Rem_right(*pos)
		*pos = nil
		bs.size --
	}
	return
}

//将两棵二叉树合并为单棵二叉树
func Merge(bs, bs2 *BiTree, data interface{}) (*BiTree, error) {
	rel := &BiTree{size: 0}
	if e := rel.Ins_left(nil, data); e != nil {
		return nil, e
	}
	rel.root.left = bs.root
	rel.root.right = bs2.root
	rel.size = rel.size + bs.size + bs2.size

	bs.size = 0
	bs.root = nil
	bs2.size = 0
	bs2.root = nil
	return rel, nil
}

func (bs *BiTree) Size() int {
	return bs.size
}

func (bs *BiTree) Root() *BiTreeNode {
	return bs.root
}


//判断node是不是二叉树的叶子结点
func (node *BiTreeNode) Is_leaf() bool {
	return node.left == nil && node.right == nil
}

//判断node的终止 end of bitree
func (node *BiTreeNode) Is_eob() bool {
	return node == nil
}

func (node *BiTreeNode) Right() *BiTreeNode {
	return node.right
}


func (node *BiTreeNode) Left() *BiTreeNode {
	return node.left
}

func (node *BiTreeNode) Data() interface{} {
	return node.data
}
