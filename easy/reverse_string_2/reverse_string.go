package reversestring2

func reverseStr(s string, k int) string {
	runes := []rune(s)
	for i := 0; i < len(runes); i += 2 * k {
		p1, p2 := i, i+k-1
		if p2 >= len(runes) {
			p2 = len(runes) - 1
		}
		for p1 < p2 {
			runes[p1], runes[p2] = runes[p2], runes[p1]
			p1++
			p2--
		}
	}

	return string(runes)
}
