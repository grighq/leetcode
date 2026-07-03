package shortestcompletingword

import "unicode"

func shortestCompletingWord(licensePlate string, words []string) string {
	countLPLetters := [26]int{}
	for _, char := range licensePlate {
		if unicode.IsLetter(char) {
			countLPLetters[unicode.ToLower(char)-'a']++
		}
	}

	res := ""
	for _, word := range words {
		countWordLetters := [26]int{}
		for _, letter := range word {
			countWordLetters[letter-'a']++
		}

		isValid := true
		for i := range 26 {
			if countLPLetters[i] > countWordLetters[i] {
				isValid = false
				break
			}
		}

		if isValid {
			if res == "" || len(res) > len(word) {
				res = word
			}
		}
	}

	return res
}
