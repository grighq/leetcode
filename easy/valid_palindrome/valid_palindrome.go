package validpalindrome

import "unicode"

func isPalindrome(s string) bool {
	runes := []rune(s)
	p1, p2 := 0, len(runes)-1

	for p1 < p2 {
		first, last := runes[p1], runes[p2]
		switch {
		case !unicode.IsLetter(first) && !unicode.IsDigit(first):
			p1++
		case !unicode.IsLetter(last) && !unicode.IsDigit(last):
			p2--
		case unicode.ToLower(first) != unicode.ToLower(last):
			return false
		default:
			p1++
			p2--
		}
	}

	return true
}
