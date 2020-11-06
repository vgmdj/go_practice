package Word_Ladder

import "math"

func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordId := make(map[string]int)
	graph := make([][]int, 0)
	addWord := func(word string) int {
		id, has := wordId[word]
		if has {
			return id
		}

		wordId[word] = len(wordId)
		graph = append(graph, []int{})

		return wordId[word]

	}

	addEdge := func(word string) int {
		id1 := addWord(word)
		bts := []byte(word)
		for i := 0; i < len(word); i++ {
			bts[i] = '*'
			id2 := addWord(string(bts))
			bts[i] = word[i]

			graph[id1] = append(graph[id1], id2)
			graph[id2] = append(graph[id2], id1)

		}

		return id1

	}

	for _, word := range wordList {
		addEdge(word)
	}

	endId, endExist := wordId[endWord]
	if !endExist {
		return 0
	}
	beginId := addEdge(beginWord)

	const initDist = math.MaxInt32
	endQueue := []int{endId}
	endDist := make([]int, len(wordId))
	for i := range endDist {
		endDist[i] = initDist
	}
	endDist[endId] = 0

	beginQueue := []int{beginId}
	beginDist := make([]int, len(wordId))
	for i := range beginDist {
		beginDist[i] = initDist
	}
	beginDist[beginId] = 0

	for len(beginQueue) > 0 && len(endQueue) > 0 {
		queue := beginQueue
		beginQueue = nil
		for _, beginLast := range queue {
			if endDist[beginLast] < initDist {
				return (beginDist[beginLast]+endDist[beginLast])/2 + 1
			}

			for _, wid := range graph[beginLast] {
				if beginDist[wid] == initDist {
					beginDist[wid] = beginDist[beginLast] + 1
					beginQueue = append(beginQueue, wid)
				}
			}
		}

		queue = endQueue
		endQueue = nil
		for _, endLast := range queue {
			if beginDist[endLast] < initDist {
				return (beginDist[endLast]+endDist[endLast])/2 + 1
			}

			for _, wid := range graph[endLast] {
				if endDist[wid] == initDist {
					endDist[wid] = endDist[endLast] + 1
					endQueue = append(endQueue, wid)
				}

			}
		}

	}

	return 0
}


