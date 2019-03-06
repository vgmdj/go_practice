package Rotate_Array

func rotate(nums []int, k int) {
	k = k % len(nums)
	if k < 1 {
		return
	}

	//first reverse all
	reverse(nums)

	//then reverse first k
	reverse(nums[:k])

	//finally reverse the last
	reverse(nums[k:])

}

func reverse(nums []int) {
	if len(nums) < 2 {
		return
	}

	var h, t = 0, len(nums) - 1
	for h < t {
		nums[h], nums[t] = nums[t], nums[h]
		h++
		t--
	}

	return

}
