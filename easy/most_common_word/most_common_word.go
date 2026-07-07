package mostcommonword

import "unicode"

func mostCommonWord(paragraph string, banned []string) string {
	wordFrequency := make(map[string]int)
	bannedWords := make(map[string]struct{}, len(banned))

	for _, word := range banned {
		bannedWords[word] = struct{}{}
	}

	runes := []rune{}
	for _, char := range paragraph + " " {
		if unicode.IsLetter(char) {
			runes = append(runes, unicode.ToLower(char))
		} else if len(runes) > 0 {
			word := string(runes)
			if _, isBanned := bannedWords[word]; !isBanned {
				wordFrequency[word]++
			}
			runes = []rune{}
		}
	}

	res, maxFrequency := "", 0
	for word, freq := range wordFrequency {
		if freq > maxFrequency {
			res = word
			maxFrequency = freq
		}
	}

	return res
}
