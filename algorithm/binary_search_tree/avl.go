package binary_search_tree


/*
二叉搜索树 ： 从根结点向下查询，若是结点值比目标值小，顺着左子树继续查找，反之顺着右子树继续查找。

插入：也是按照规则，从根结点开始，一直查询到空位置进行插入

复杂度 o(lg n) : n代表树中结点的个数 （平衡二叉树的情况）



AVL 树 ：特殊类型的二叉树，每个结点保存一份额外的信息：结点的平衡因子.结点左子树的高度减去右子树的高度。
插入结点时avl树通过自我调整，使平衡因子始终保持在 +1，-1，0。该过程称为旋转

设 x 为刚插入AVL树的结点，设A为离x最近的且平衡因子更改为+-2的结点，


当 x 位于A的左子树的左子树上时，进行 LL 旋转
   left为A的左子树（+1），

   将A的左指针指向left的右子结点，left的右指针指向A，将原来指向A的指针指向left。 平衡因子（left ：0，A：0）


当 x 位于A的左子树的右子树上时，执行 LR 旋转
  left为A的左子树（-1） ，并设A的子孙结点grandchild为left的右子结点。

left的右子节点指向grandchild的左子结点，（left结点变为0或不变）
grandchild的左子结点指向left，
A的左子结点指向grandchild的右子结点，（A变为0或-1）（此时A的层数于left层数相同）
grandchild的右子结点指向A，（grandchild为0）
将原指向A的指针指向grandchild。




//互为对称


当 x 位于A的右子树的左子树上时，执行 RR 旋转
   right为A的右子树（+1）

  将A的右指针指向right的左子结点，right的左指针指向A，将原来指向A的指针指向right。 平衡因子（left ：0，A：0）


当 x 位于A的右子树的右子树上时，执行 RL 旋转
  right为A的左子树（-1） ，并设A的子孙结点grandchild为right的左子结点。

right的左子节点指向grandchild的右子结点，（left结点变为0或不变）
grandchild的右子结点指向right，
A的右子结点指向grandchild的左子结点，（A变为0或-1）（此时A的层数于left层数相同）
grandchild的左子结点指向A，（grandchild为0）
将原指向A的指针指向grandchild。

*/

const (
	AVL_LEFT_HEAVY  = 1
	AVL_BALANCE     = 0
	AVL_RIGHT_HEAVY = -1
)

type BisTree BiTree

type AvlNode struct {
	data interface{}
	//用来标识结点是否已经处于删除状态 false 表示存在 true 表示已删除
	hidden bool
	//该结点的平衡因子
	factor int
}

//BiTreeNode的data为AvlNode结点

//node为指向A的指针 注意是**
func Rorate_left(node **BiTreeNode) {
	var left, grandchild *BiTreeNode

	left = (*node).left

	//LL
	if (left.data.(*AvlNode).factor == AVL_LEFT_HEAVY) {
		(*node).left = left.right
		left.right = *node

		left.data.(*AvlNode).factor = AVL_BALANCE
		(*node).data.(*AvlNode).factor = AVL_BALANCE
		*node = left
	} else { //LR
		grandchild = left.right
		left.right = grandchild.left
		grandchild.left = left
		(*node).left = grandchild.right
		grandchild.right = (*node)

		switch grandchild.data.(*AvlNode).factor {
		case AVL_LEFT_HEAVY:
			left.data.(*AvlNode).factor = AVL_BALANCE
			(*node).data.(*AvlNode).factor = AVL_RIGHT_HEAVY
		case AVL_BALANCE:
			left.data.(*AvlNode).factor = AVL_BALANCE
			(*node).data.(*AvlNode).factor = AVL_BALANCE
		case AVL_RIGHT_HEAVY:
			left.data.(*AvlNode).factor = AVL_LEFT_HEAVY
			(*node).data.(*AvlNode).factor = AVL_BALANCE
		}

		grandchild.data.(*AvlNode).factor = AVL_BALANCE
		*node = grandchild
	}
}

//互为镜像: node为指向A的指针 注意是**
func Rorate_right(node **BiTreeNode) {
	var right, grandchild *BiTreeNode

	right = (*node).right

	//RR
	if (right.data.(*AvlNode).factor == AVL_RIGHT_HEAVY) {
		(*node).right = right.left
		right.right = *node

		right.data.(*AvlNode).factor = AVL_BALANCE
		(*node).data.(*AvlNode).factor = AVL_BALANCE
		*node = right
	} else { //RL
		grandchild = right.right
		right.left = grandchild.right
		grandchild.right = right
		(*node).right = grandchild.left
		grandchild.left = (*node)

		switch grandchild.data.(*AvlNode).factor {
		case AVL_LEFT_HEAVY:
			right.data.(*AvlNode).factor = AVL_RIGHT_HEAVY
			(*node).data.(*AvlNode).factor = AVL_BALANCE
		case AVL_BALANCE:
			right.data.(*AvlNode).factor = AVL_BALANCE
			(*node).data.(*AvlNode).factor = AVL_BALANCE
		case AVL_RIGHT_HEAVY:
			right.data.(*AvlNode).factor = AVL_BALANCE
			(*node).data.(*AvlNode).factor = AVL_LEFT_HEAVY
		}

		grandchild.data.(*AvlNode).factor = AVL_BALANCE
		*node = grandchild
	}
}

//销毁某个结点下方的左子树
func Destory_left(bs *BisTree, node *BiTreeNode) {
	var pos *BiTreeNode

	if bs.size == 0 {
		return
	}

	if node == nil {
		pos = bs.root
	} else {
		pos = node.left
	}

	if pos != nil {
		Destory_left(bs, pos)
		Destory_right(bs, pos)
		bs.size--
	}
}

//销毁某个结点下方的右子树
func Destory_right(bs *BisTree, node *BiTreeNode) {
	var pos **BiTreeNode

	if bs.size == 0 {
		return
	}

	if node == nil {
		pos = &bs.root
	} else {
		pos = &node.right
	}

	if *pos != nil {
		Destory_left(bs, *pos)
		Destory_right(bs, *pos)
		*pos = nil
		bs.size--
	}
}

//AVL树的插入 node为A结点
func Insert(tree *BiTree, node **BiTreeNode, data interface{}, balanced *bool) {
	var avlNode *AvlNode

	if (*node).Is_eob() {
		avlNode = new(AvlNode)
		avlNode.data = data
		avlNode.factor = AVL_BALANCE
		avlNode.hidden = false
		tree.Ins_left(*node, avlNode)
		return
	} else {

		var cmpval int

		//值大于当前结点值
		cmpval = tree.compare(data, (*node).data.(*AvlNode).data)

		//放入到左子树中
		if cmpval < 0 {
			if (*node).left.Is_eob() {
				avlNode = new(AvlNode)
				avlNode.data = data
				avlNode.factor = AVL_BALANCE
				avlNode.hidden = false
				tree.Ins_left(*node, avlNode)
				*balanced = false
			} else {
				//以左结点为根做递归插入
				Insert(tree, &(*node).left, data, balanced)
			}

			if !(*balanced) {
				switch (*node).data.(*AvlNode).factor {
				case AVL_LEFT_HEAVY:
					Rorate_left(node)
					*balanced = true
				case AVL_BALANCE:
					(*node).data.(*AvlNode).factor = AVL_LEFT_HEAVY
				case AVL_RIGHT_HEAVY:
					(*node).data.(*AvlNode).factor = AVL_BALANCE
					*balanced = true
				}

			}

		} else if cmpval > 0 { //插入到右子树中
			if (*node).right.Is_eob() {
				avlNode = new(AvlNode)
				avlNode.data = data
				avlNode.factor = AVL_BALANCE
				avlNode.hidden = false
				tree.Ins_right(*node, avlNode)
				*balanced = false
			} else {
				//以左结点为根做递归插入
				Insert(tree, &(*node).right, data, balanced)
			}

			if !(*balanced) {
				switch (*node).data.(*AvlNode).factor {
				case AVL_LEFT_HEAVY:
					(*node).data.(*AvlNode).factor = AVL_BALANCE
					*balanced = true
				case AVL_BALANCE:
					(*node).data.(*AvlNode).factor = AVL_RIGHT_HEAVY
				case AVL_RIGHT_HEAVY:
					Rorate_right(node)
					*balanced = true
				}
			}
		} else {
			//数据存在
			if !(*node).data.(*AvlNode).hidden {
				return
			} else { //数据状态修改为为删除
				(*node).data.(*AvlNode).hidden = false
				*balanced = true
			}
		}
	}
}

//从node开始查找 删除结点为data的结点
func Hide(tree *BisTree, node *BiTreeNode, data interface{}) {
	var cmpval int
	if node.Is_eob() {
		return
	}

	cmpval = tree.compare(data, (*node).data.(*AvlNode).data)

	if cmpval < 0 {
		Hide(tree, node.left, data)
	} else if cmpval > 0 {
		Hide(tree, node.right, data)
	} else {
		(*node).data.(*AvlNode).hidden = true
	}

	return
}

func Lookup(tree *BisTree, node *BiTreeNode, data interface{}) int {
	var cmpval, retval int
	if node.Is_eob() {
		return -1
	}

	cmpval = tree.compare(data, (*node).data.(*AvlNode).data)

	if cmpval < 0 {
		retval = Lookup(tree, node.left, data)
	} else if cmpval > 0 {
		retval = Lookup(tree, node.right, data)
	} else {
		if !(*node).data.(*AvlNode).hidden {
			data = (*node).data.(*AvlNode).data
			retval = 0
		} else {
			return -1
		}
	}

	return retval
}
