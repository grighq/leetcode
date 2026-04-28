package findindexofthefirstoccurrenceinastring

func strStr(haystack, needle string) int {
	for i := 0; len(needle) <= len(haystack[i:]); i++ {
		if needle == haystack[i:i+len(needle)] {
			return i
		}
	}
	return -1
}
