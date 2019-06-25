package Binary_Tree__Zigzag_Level_Order_Traversal

import "container/list"

/*
use s1 to save the level nodes
use s2 to save the level+1 nodes
then swap s1 and s2

example: tree [1,2,3,4,5,6,7,8,9,10]
if level is odd,  like 1
s1(level 1):                s2(level2):
                                  7
					              6
3                                 5
2                                 4

if level is even , like 2:
s1(level 2, use level 1's s2):                        s2(level3):
       7                                                   10
       6                                                    9
       5                                                    8
       4

*/

func zigzagLevelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)

	level := 0
	s1, s2 := NewStack(), NewStack()
	s1.Push(root)

	for !s1.IsEmpty() {
		for !s1.IsEmpty() {
			elem := s1.Pop()
			if elem == nil {
				continue
			}

			if len(result) <= level {
				result = append(result, []int{elem.Val})

			} else {
				result[level] = append(result[level], elem.Val)

			}

			if level%2 == 0 {
				s2.Push(elem.Left)
				s2.Push(elem.Right)

			} else {
				s2.Push(elem.Right)
				s2.Push(elem.Left)

			}

		}

		level++

		s1, s2 = s2, s1

	}

	return result

}

type Stack struct {
	l *list.List
}

func NewStack() *Stack {
	return &Stack{
		l: list.New(),
	}
}

func (s *Stack) Push(node *TreeNode) {
	s.l.PushBack(node)
}

func (s *Stack) Pop() *TreeNode {
	if s.l.Len() == 0 {
		return nil
	}

	e := s.l.Back()
	s.l.Remove(e)
	return e.Value.(*TreeNode)
}

func (s *Stack) IsEmpty() bool {
	return s.l.Len() == 0
}
