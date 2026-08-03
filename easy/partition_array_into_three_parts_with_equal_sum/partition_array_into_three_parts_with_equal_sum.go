package partitionarrayintothreepartswithequalsum

func canThreePartsEqualSum(arr []int) bool {
	sum := 0
	for _, num := range arr {
		sum += num
	}

	if sum%3 != 0 {
		return false
	}

	target := sum / 3

	count := 0
	currSum := 0
	for _, num := range arr[:len(arr)-1] {
		currSum += num

		if currSum == target {
			count++
			currSum = 0
		}
	}

	return count >= 2
}
