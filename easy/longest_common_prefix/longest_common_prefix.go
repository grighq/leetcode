package longestcommonprefix

import (
	"slices"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	slices.Sort(strs)
	prefix := strs[0]
	for prefix != "" {
		if strings.Index(strs[len(strs)-1], prefix) == 0 {
			return prefix
		}
		prefix = prefix[:len(prefix)-1]
	}
	return prefix
}
