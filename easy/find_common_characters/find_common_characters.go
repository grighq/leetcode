package findcommoncharacters

func commonChars(words []string) []string {
	charsFreq := [26]int{}
	for _, char := range []byte(words[0]) {
		charsFreq[char-'a']++
	}

	for _, word := range words[1:] {
		currFreq := [26]int{}
		for _, char := range []byte(word) {
			currFreq[char-'a']++
		}

		for i := range 26 {
			charsFreq[i] = min(charsFreq[i], currFreq[i])
		}
	}

	charsCount := 0
	for _, count := range charsFreq {
		charsCount += count
	}

	idx := 0
	res := make([]string, charsCount)
	for i := range charsFreq {
		for charsFreq[i] > 0 {
			res[idx] = string(byte(i + 'a'))
			idx++
			charsFreq[i]--
		}
	}

	return res
}
