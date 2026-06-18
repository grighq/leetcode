package canplaceflowers

func canPlaceFlowers(flowerbed []int, n int) bool {
	freePlots := 0
	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] == 1 {
			i++
		} else if i+1 == len(flowerbed) || flowerbed[i+1] != 1 {
			freePlots++
			i++
		}
	}

	return freePlots >= n
}
