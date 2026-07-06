package uniquemorsecodewords

import "strings"

func uniqueMorseRepresentations(words []string) int {
	res := make(map[string]struct{})
	morse := [26]string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}

	for _, word := range words {
		var morseWord strings.Builder
		for _, letter := range word {
			morseWord.WriteString(morse[letter-'a'])
		}

		res[morseWord.String()] = struct{}{}
	}

	return len(res)
}
