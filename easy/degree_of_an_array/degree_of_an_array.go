package degreeofanarray

type degree struct {
	count, firstIdx, lastIdx int
}

func findShortestSubArray(nums []int) int {
	degrees := make(map[int]*degree, len(nums))
	for i, num := range nums {
		if d, ok := degrees[num]; !ok {
			degrees[num] = &degree{1, i, i}
		} else {
			d.count++
			d.lastIdx = i
		}
	}

	maxDegrees, res := 0, 50001
	for _, d := range degrees {
		if d.count < maxDegrees {
			continue
		}

		if d.count > maxDegrees {
			maxDegrees = d.count
			res = d.lastIdx - d.firstIdx + 1
		} else {
			res = min(res, d.lastIdx-d.firstIdx+1)
		}

	}

	return res
}
