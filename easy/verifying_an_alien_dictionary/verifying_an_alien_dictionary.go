package verifyinganaliendictionary

func isAlienSorted(words []string, order string) bool {
	charsWeight := [26]int{}
	for i := range order {
		charsWeight[order[i]-'a'] = i
	}

	for i := 1; i < len(words); i++ {
		currWord, nextWord := words[i-1], words[i]

		j := 0
		for ; j < len(currWord) && j < len(nextWord); j++ {
			currChar, nextChar := currWord[j], nextWord[j]
			currCharWeight, nextCharWeight := charsWeight[currChar-'a'], charsWeight[nextChar-'a']

			if currCharWeight < nextCharWeight {
				break
			} else if currCharWeight > nextCharWeight {
				return false
			}
		}

		if j == len(nextWord) && len(currWord) > len(nextWord) {
			return false
		}
	}

	return true
}
