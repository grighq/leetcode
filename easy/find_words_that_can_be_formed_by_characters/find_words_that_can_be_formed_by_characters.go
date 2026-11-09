package findwordsthatcanbeformedbycharacters

func countCharacters(words []string, chars string) int {
	countChars := [26]int{}
	for _, char := range []byte(chars) {
		countChars[char-'a']++
	}

	res := 0
	for _, word := range words {
		isGood := true
		currChars := countChars

		for _, letter := range []byte(word) {
			idx := letter - 'a'
			currChars[idx]--

			if currChars[idx] < 0 {
				isGood = false
				break
			}
		}

		if isGood {
			res += len(word)
		}
	}

	return res
}
