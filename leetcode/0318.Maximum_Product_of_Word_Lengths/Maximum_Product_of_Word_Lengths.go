package Maximum_Product_of_Word_Lengths

func maxProduct(words []string) int {
	bitMap := make([]int32, len(words))
	max := 0

	for k, v := range words {
		for _, bv := range v {
			bitMap[k] = bitMap[k] | (1 << uint(bv-'a'))
		}
	}

	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if bitMap[i]|bitMap[j] != 0 {
				continue
			}

			if max < len(words[i])*len(words[j]) {
				max = len(words[i]) * len(words[j])
			}
		}

	}

	return max

}
