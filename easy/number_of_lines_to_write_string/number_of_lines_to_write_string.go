package numberoflinestowritestring

func numberOfLines(widths []int, s string) []int {
	width, countLines := 0, 1
	for _, letter := range s {
		w := widths[letter-'a']
		if width > 100-w {
			width = w
			countLines++
		} else {
			width += w
		}
	}

	return []int{countLines, width}
}
