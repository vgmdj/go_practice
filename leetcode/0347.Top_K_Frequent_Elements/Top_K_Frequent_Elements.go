package Top_K_Frequent_Elements

import "container/heap"

type Item struct {
	value    int
	priority int
	index    int
}
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

// min heap
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].priority < pq[j].priority }

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func topKFrequent(nums []int, k int) []int {
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	frequent := make(map[int]int)
	for _, v := range nums {
		frequent[v] += 1
	}
	for key, v := range frequent {
		heap.Push(&pq, &Item{value: key, priority: v})
		if len(pq) > k {
			heap.Pop(&pq)
		}
	}
	res := make([]int, k)
	for i, j := k-1, 0; i >= 0; i, j = i-1, j+1 {
		res[j] = pq[i].value
	}
	return res
}

func topKFrequent2(nums []int, k int) []int {
	counter := make(map[int]int, k)
	for _, val := range nums {
		counter[val]++
	}

	results := make([]int, 0, k)
	for k > 0 {
		maxk, maxv := 0, 0
		for key, val := range counter {
			if maxv < val {
				maxv = val
				maxk = key
			}
		}

		results = append(results, maxk)
		delete(counter, maxk)
		k--

	}

	return results

}

func topKFrequent3(nums []int, k int) []int {
	var rst []int
	numsMap := map[int]int{}
	for _, v := range nums {
		numsMap[v]++
	}

	numsMapSort := map[int][]int{}
	maxTemp := make([]int, len(nums)+1)
	for k, v := range numsMap {
		numsMapSort[v] = append(numsMapSort[v], k)
		maxTemp[v] = 1
	}

	for i := len(maxTemp) - 1; i >= 0; i-- {
		if maxTemp[i] != 0 {
			for _, v := range numsMapSort[i] {
				if k == 0 {
					return rst
				}
				rst = append(rst, v)
				k--
			}
		}
	}
	return rst
}
