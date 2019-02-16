package _017_Letter_Combinations_of_a_Phone_Number

func letterCombinations(digits string) []string {
	var result []string

	var reflect = map[int32][]string{
		'2': []string{"a", "b", "c"},
		'3': []string{"d", "e", "f"},
		'4': []string{"g", "h", "i"},
		'5': []string{"j", "k", "l"},
		'6': []string{"m", "n", "o"},
		'7': []string{"p", "q", "r", "s"},
		'8': []string{"t", "u", "v"},
		'9': []string{"w", "x", "y", "z"},
	}

	for _, v := range digits {

		for resk, resv := range result {
			for refk,refv := range reflect[v] {
				if refk == 0{
					result[resk] +=  refv
					continue
				}
				result = append(result, resv+refv)
			}
		}

		if len(result) == 0 {
			for _, refv := range reflect[v] {
				result = append(result, refv)
			}
		}

	}

	return result

}
