package assigncookies

import (
	"slices"
)

func findContentChildren(g, s []int) int {
	slices.Sort(g)
	slices.Sort(s)
	p1, p2 := 0, 0
	for p1 < len(g) && p2 < len(s) {
		if g[p1] <= s[p2] {
			p1++
		}
		p2++
	}

	return p1
}
