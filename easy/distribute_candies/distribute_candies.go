package distributecandies

func distributeCandies(candyType []int) int {
	candies := make(map[int]struct{})

	for _, candy := range candyType {
		candies[candy] = struct{}{}
	}

	if len(candyType)/2 >= len(candies) {
		return len(candies)
	}

	return len(candyType) / 2
}

// func distributeCandies(candyType []int) int {
// 	slices.Sort(candyType)
// 	count, currCandy := 1, candyType[0]
//
// 	for _, candy := range candyType {
// 		if candy != currCandy {
// 			count++
// 			currCandy = candy
// 		}
// 	}
//
// 	if len(candyType)/2 >= count {
// 		return count
// 	}
//
// 	return len(candyType) / 2
// }
