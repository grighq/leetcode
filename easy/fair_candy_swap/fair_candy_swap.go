package faircandyswap

func fairCandySwap(aliceSizes []int, bobSizes []int) []int {
	aliceSum, bobSum := 0, 0
	aliceCandies := make(map[int]struct{})

	for _, size := range aliceSizes {
		aliceSum += size
		aliceCandies[size] = struct{}{}
	}

	for _, size := range bobSizes {
		bobSum += size
	}

	res := make([]int, 2)
	delta := (aliceSum - bobSum) / 2
	for _, size := range bobSizes {
		if _, ok := aliceCandies[size+delta]; ok {
			res[0], res[1] = size+delta, size
			break
		}
	}

	return res
}
