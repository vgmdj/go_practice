package Majority_Element_II

func majorityElement(nums []int) []int {
	element1, element2 := 0, 0
	vote1, vote2 := 0, 0
	for _, num := range nums {
		if element1 == num {
			vote1++

		} else if element2 == num {
			vote2++

		} else if vote1 == 0 {
			element1 = num
			vote1 = 1

		} else if vote2 == 0 {
			element2 = num
			vote2 = 1

		} else {
			vote1--
			vote2--

		}
	}

	cnt1, cnt2 := 0, 0
	for _, num := range nums {
		if element1 == num {
			cnt1++

		} else if element2 == num {
			cnt2++

		}
	}

	result := make([]int, 0)
	if cnt1 > len(nums)/3 {
		result = append(result, element1)
	}

	if cnt2 > len(nums)/3 {
		result = append(result, element2)
	}

	return result

}
