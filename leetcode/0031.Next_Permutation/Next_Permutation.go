package Next_Permutation

func nextPermutation(nums []int) {
	if len(nums) < 2 {
		return
	}

	p := len(nums) - 1
	for ; p > 0; p-- {
		if nums[p] > nums[p-1] {
			break
		}
	}

	for q := p; q < len(nums); q++ {
		if p > 0 && nums[q] > nums[p-1] && (q == len(nums)-1 || nums[q+1] <= nums[p-1]) {
			nums[q], nums[p-1] = nums[p-1], nums[q]
			break
		}
	}

	for q := len(nums) - 1; p < q; {
		nums[p], nums[q] = nums[q], nums[p]
		p++
		q--
	}

}
