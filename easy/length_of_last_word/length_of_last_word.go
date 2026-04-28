package lengthoflastword

func lengthOfLastWord(s string) int {
	count := 0
	for end := len(s) - 1; end >= 0; end-- {
		if s[end] != ' ' {
			count++
		} else if count > 0 {
			return count
		}
	}
	return count
}
