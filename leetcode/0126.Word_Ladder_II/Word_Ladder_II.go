package Word_Ladder_II

import "math"

func findLadders(beginWord string, endWord string, wordList []string) [][]string {

	wordID := make(map[string]int, len(wordList)+1)
	for i, c := range wordList {
		wordID[c] = i
	}

	if _, ok := wordID[endWord]; !ok {
		return nil
	}

	if _, ok := wordID[beginWord]; !ok {
		wordList = append(wordList, beginWord)
		wordID[beginWord] = len(wordList) - 1
	}

	edges := make([][]int, len(wordList))
	for i := 0; i < len(wordList); i++ {
		for j := i + 1; j < len(wordList); j++ {
			if canTransform(wordList[i], wordList[j]) {
				edges[i] = append(edges[i], j)
				edges[j] = append(edges[j], i)
			}
		}
	}

	result := make([][]string, 0)
	addResult := func(wordIndexList []int) {
		words := make([]string, 0)
		for _, wordIndex := range wordIndexList {
			words = append(words, wordList[wordIndex])
		}
		result = append(result, words)
	}

	dist := make([]int, len(wordList))
	for i := range dist {
		dist[i] = math.MaxInt32
	}
	dist[wordID[beginWord]] = 0

	queue := [][]int{{wordID[beginWord]}}

	for len(queue) != 0 {
		bottom := queue[0]
		lastWordID := bottom[len(bottom)-1]
		queue = queue[1:]

		if lastWordID == wordID[endWord] {
			addResult(bottom)
		}

		for _, next := range edges[lastWordID] {
			// not shortest
			if dist[lastWordID]+1 > dist[next] {
				continue
			}

			// can visit
			dist[next] = dist[lastWordID] + 1
			path := make([]int, len(bottom)+1)
			copy(path, bottom)
			path[len(path)-1] = next

			queue = append(queue, path)

		}

	}

	return result

}

func canTransform(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return a[i+1:] == b[i+1:]
		}
	}

	return true

}
