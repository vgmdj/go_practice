package Substring_with_Concatenation_of_All_Words

func findSubstring(s string, words []string) []int {
	if len(words) == 0 {
		return []int{}
	}

	length := len(words[0])
	size := len(words) * length
	wordMap := make(map[string]int, len(words))
	for _, v := range words {
		wordMap[v]++
	}

	result := make([]int, 0)
	for i := 0; i <= len(s)-size; i++ {
		if match(s[i:i+size], wordMap, length) {
			result = append(result, i)
		}
	}

	return result
}

func match(text string, wordMap map[string]int, length int) bool {
	counter := make(map[string]int, len(wordMap))
	for i := 0; i <= len(text)-length; i += length {
		str := text[i : i+length]
		counter[str]++

		if counter[str] > wordMap[str] {
			return false
		}
	}

	return true

}
