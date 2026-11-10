package uniquenumberofoccurrences

func uniqueOccurrences(arr []int) bool {
	freq := [2001]int{}
	for _, num := range arr {
		freq[num+1000]++
	}

	uniqFreq := [1001]bool{}
	for _, count := range freq {
		if count > 0 && uniqFreq[count] {
			return false
		} else {
			uniqFreq[count] = true
		}
	}

	return true
}
