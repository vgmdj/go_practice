package Next_Greater_Element_I

import "container/list"

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	stack := list.New()
	greater := make(map[int]int, len(nums2))

	for _, v := range nums2 {

		for stack.Len() != 0 {
			e := stack.Back()

			if e.Value.(int) >= v {
				break

			}

			stack.Remove(e)
			greater[e.Value.(int)] = v

		}

		stack.PushBack(v)

	}

	for k, v := range nums1 {
		n, ok := greater[v]
		if !ok {
			nums1[k] = -1
			continue
		}

		nums1[k] = n

	}

	return nums1

}
