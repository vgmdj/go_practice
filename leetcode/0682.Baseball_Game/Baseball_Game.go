package Baseball_Game

import (
	"container/list"
	"strconv"
)

func calPoints(ops []string) int {
	result := 0
	stack := list.New()

	for _, v := range ops {
		switch v {
		default:
			p, _ := strconv.Atoi(v)
			stack.PushBack(p)

		case "+":
			p2 := stack.Back()
			stack.Remove(p2)

			p1 := stack.Back()

			stack.PushBack(p2.Value.(int))
			stack.PushBack(p1.Value.(int) + p2.Value.(int))

		case "D":
			p := stack.Back().Value.(int)
			stack.PushBack(p * 2)

		case "C":
			stack.Remove(stack.Back())

		}

	}

	for stack.Len() != 0 {
		e := stack.Back()
		stack.Remove(e)

		result += e.Value.(int)

	}

	return result

}
