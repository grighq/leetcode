package reversewordsinastring3

import "unicode"

func reverseWords(s string) string {
	runes := []rune(s)
	for p1, p2 := 0, 0; p2 <= len(runes); p2++ {
		if p2 == len(runes) || unicode.IsSpace(runes[p2]) {
			l, r := p1, p2-1
			for l < r {
				runes[l], runes[r] = runes[r], runes[l]
				l++
				r--
			}
			p1 = p2 + 1
		}
	}

	return string(runes)
}
