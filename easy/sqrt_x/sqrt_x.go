package sqrtx

func mySqrt(x int) int {
	if x < 2 {
		return x
	}

	start, end := 2, x/2
	for start <= end {
		mid := start + (end-start)/2
		if mid > x/mid {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return end
}
