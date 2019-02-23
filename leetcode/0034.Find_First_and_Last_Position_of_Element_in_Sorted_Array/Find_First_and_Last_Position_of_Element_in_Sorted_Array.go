package Find_First_and_Last_Position_of_Element_in_Sorted_Array

func searchRange(nums []int, target int) []int {
	var start, end = -1, -1
	var left, right, mid = 0, len(nums) - 1, 0

	//先圈定target的大致范围
	for left <= right {
		mid = left + (right-left)/2

		if nums[mid] == target {
			start = mid
			end = mid
			break

		} else if target < nums[mid] {
			right = mid - 1

		} else {
			left = mid + 1

		}

	}

	//说明不存在
	if start == -1 && end == -1 {
		return []int{start, end}
	}

	//再找start位置
	for left <= start {
		mid = left + (start-left)/2
		if nums[mid] == target {
			start = mid - 1

		} else {
			left = mid + 1
		}

	}

	//最后确定end位置
	for end <= right {
		mid = end + (right-end)/2
		if nums[mid] == target {
			end = mid + 1

		} else {
			right = mid - 1

		}

	}

	return []int{start+1, end-1}

}
