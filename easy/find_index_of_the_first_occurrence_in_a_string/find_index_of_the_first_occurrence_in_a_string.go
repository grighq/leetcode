package findindexofthefirstoccurrenceinastring

func strStr(haystack, needle string) int {
	for k := len(needle); k <= len(haystack); k++ {
		if haystack[k-len(needle):k] == needle {
			return k - len(needle)
		}
	}
	return -1
}
