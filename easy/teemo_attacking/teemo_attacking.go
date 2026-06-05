package teemoattacking

func findPoisonedDuration(timeSeries []int, duration int) int {
	count := 0
	for i, t := range timeSeries[1:] {
		interval := t - timeSeries[i]
		count += min(duration, interval)
	}

	return count + duration
}
