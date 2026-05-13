package maximumrepeatingsubstring

func maxRepeating(sequence, word string) int {
	k := 0
	dp := make([]int, len(sequence))
	for i := len(word) - 1; i < len(sequence); i++ {
		if word == sequence[i-len(word)+1:i+1] {
			if i-len(word) >= 0 {
				dp[i] = dp[i-len(word)] + 1
			} else {
				dp[i] = 1
			}
		}
		k = max(k, dp[i])
	}
	return k
}

// func maxRepeating(sequence, word string) int {
// 	k, count := 0, 0
// 	for i := 0; i+len(word) <= len(sequence); {
// 		if word == sequence[i:i+len(word)] {
// 			count++
// 			i += len(word)
// 		} else if count != 0 {
// 			k = max(k, count)
// 			i = i - len(word)*count + 1
// 			count = 0
// 		} else {
// 			i++
// 		}
// 	}
//
// 	return max(k, count)
// }
