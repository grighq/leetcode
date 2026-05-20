package reversevowelsofastring

import (
	"strings"
)

func reverseVowels(s string) string {
	runes := []rune(s)
	vowels := "aeiouAEIOU"
	p1, p2 := 0, len(runes)-1

	for p1 < p2 {
		if !strings.ContainsRune(vowels, runes[p1]) {
			p1++
		} else if !strings.ContainsRune(vowels, runes[p2]) {
			p2--

		} else {
			runes[p1], runes[p2] = runes[p2], runes[p1]
			p1++
			p2--
		}
	}

	return string(runes)
}
