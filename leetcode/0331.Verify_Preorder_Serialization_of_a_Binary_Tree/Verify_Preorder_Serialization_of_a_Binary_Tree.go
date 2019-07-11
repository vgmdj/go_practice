package Verify_Preorder_Serialization_of_a_Binary_Tree

import (
	"container/list"
	"strings"
)

func isValidSerialization(preorder string) bool {
	trees := strings.Split(preorder, ",")
	stack := list.New()
	stack.PushBack("#")
	for _, v := range trees {
		if v != "#" {
			stack.PushBack(v)
			continue
		}

		if stack.Len() == 0 || (stack.Len() == 1 && stack.Back().Value.(string) != "#") {
			return false
		}

		stack.Remove(stack.Back())

	}

	if stack.Len() != 0 {
		return false
	}

	return true

}
