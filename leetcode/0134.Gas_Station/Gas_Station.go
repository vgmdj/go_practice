package Gas_Station

func canCompleteCircuit(gas []int, cost []int) int {
	start, remain, debt := 0, 0, 0

	for k, v := range gas {
		remain += v - cost[k]
		if remain < 0 {
			start = k+1
			debt += remain
			remain = 0

		}
	}

	if remain+debt < 0 {
		return -1
	}

	return start

}
