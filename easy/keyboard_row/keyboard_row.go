package keyboardrow

import (
	"unicode"
)

func findWords(words []string) []string {
	res := []string{}
	charMap := make(map[rune]int, 26)
	keyboardRows := []string{"qwertyuiop", "asdfghjkl", "zxcvbnm"}

	for rowNum, row := range keyboardRows {
		for _, char := range row {
			charMap[char] = rowNum
		}
	}

	for _, word := range words {
		runes := []rune(word)
		rowNum := charMap[unicode.ToLower(runes[0])]
		isValid := true

		for _, char := range runes[1:] {
			if charMap[unicode.ToLower(char)] != rowNum {
				isValid = false
				break
			}
		}

		if isValid {
			res = append(res, word)
		}
	}

	return res
}
