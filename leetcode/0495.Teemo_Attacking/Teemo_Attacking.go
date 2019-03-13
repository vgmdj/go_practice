package Teemo_Attacking

func findPoisonedDuration(timeSeries []int, duration int) int {
	if len(timeSeries) < 1 {
		return 0
	}

	var sum int
	for i := 1; i < len(timeSeries); i++ {
		if timeSeries[i]-timeSeries[i-1] > duration {
			sum += duration
			continue
		}

		sum += timeSeries[i] - timeSeries[i-1]

	}

	return sum + duration
}
