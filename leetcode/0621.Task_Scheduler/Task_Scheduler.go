package Task_Scheduler

import "sort"

func leastInterval(tasks []byte, n int) int {
	if n == 0 {
		return len(tasks)
	}

	bucket := make([]int, 26)
	for _, c := range tasks {
		bucket[c-'A']++
	}

	sort.Ints(bucket)

	res, lastCount := 0, 1
	for i := len(bucket) - 1; i >= 0 && bucket[i] > 0; i-- {
		interval := (bucket[i]-1)*(n+1) + lastCount
		if interval <= res {
			break
		}

		res = interval
		lastCount++
	}

	if res < len(tasks) {
		return len(tasks)
	}

	return res
}
