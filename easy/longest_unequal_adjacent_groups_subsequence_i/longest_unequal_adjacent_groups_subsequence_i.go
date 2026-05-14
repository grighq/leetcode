package longestunequaladjacentgroupssubsequencei

func getLongestSubsequence(words []string, groups []int) []string {
	prev := groups[0]
	dp := []string{words[0]}
	for i, group := range groups[1:] {
		if group != prev {
			dp = append(dp, words[i+1])
			prev = group
		}
	}

	return dp
}
