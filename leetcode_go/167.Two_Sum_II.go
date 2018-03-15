package main

func twoSum(numbers []int, target int) []int {
	length := len(numbers)
	index1 := 0
	index2 := length - 1
	result := []int{}

	for index1 < index2 {
		if numbers[index1]+numbers[index2] == target {
			result = append(result, index1+1)
			result = append(result, index2+1)
			return result
		}

		if numbers[index1]+numbers[index2] > target {
			index2 = index2 - 1
		} else {
			index1 = index1 + 1
		}

	}

	return result

}
