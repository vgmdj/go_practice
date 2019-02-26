package Sort_Colors

func sortColors(nums []int) {
	var p, q, c = 0, len(nums) - 1, 0

	for c <= q {
		switch nums[c] {
		case 0:
			if c != p {
				nums[c], nums[p] = nums[p], nums[c]
			}

			c++
			p++

		case 1:
			c++

		case 2:
			nums[c], nums[q] = nums[q], nums[c]
			q--

		}

	}

}
