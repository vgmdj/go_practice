package Text_Justification

import "strings"

func fullJustify(words []string, maxWidth int) []string {
	result := make([]string, 0)
	left, right, wordsLength := 0, 0, 0
	for right < len(words) {
		wordsLength += len(words[right])
		if wordsLength+right-left <= maxWidth {
			right++

		} else {
			result = append(result, stretchAlign(words[left:right], wordsLength-len(words[right]), maxWidth))
			left = right
			wordsLength = 0
		}
	}

	if left != right {
		result = append(result, leftAlign(words[left:right], maxWidth))
	}

	return result

}

func stretchAlign(words []string, wordsLength, width int) string {
	if len(words) == 1 {
		return fillInSpace(words[0], width)
	}

	for i := 0; wordsLength < width; i++ {
		words[i%(len(words)-1)] += " "
		wordsLength++
	}
	return strings.Join(words, "")
}

func leftAlign(words []string, width int) string {
	var result string
	for _, v := range words {
		result += v + " "
	}

	if len(result) == width+1 {
		return result[:len(result)-1]
	}

	return fillInSpace(result, width)
}

func fillInSpace(word string, width int) string {
	for i := len(word); i < width; i++ {
		word += " "
	}

	return word
}
