package Find_All_Numbers_Disappeared_in_an_Array

func findDisappearedNumbers(nums []int) []int {

	for index := 0; index < len(nums); {
		if nums[index] == index+1 || nums[index] == 0 {
			index++
			continue
		}

		if nums[index] == nums[nums[index]-1] {
			nums[index] = 0
			index++
			continue
		}

		nums[nums[index]-1], nums[index] = nums[index], nums[nums[index]-1]

	}

	result := make([]int, 0)

	for k, v := range nums {
		if v == 0 {
			result = append(result, k+1)
		}
	}

	return result

}
