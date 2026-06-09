package relativeranks

import (
	"slices"
	"strconv"
)

func findRelativeRanks(score []int) []string {
	scoreIdx := make(map[int]int, len(score))
	for i, s := range score {
		scoreIdx[s] = i
	}

	slices.Sort(score)
	answer := make([]string, len(score))

	for i := len(score) - 1; i >= 0; i-- {
		idx := scoreIdx[score[i]]
		rank := len(score) - i

		switch rank {
		case 1:
			answer[idx] = "Gold Medal"
		case 2:
			answer[idx] = "Silver Medal"
		case 3:
			answer[idx] = "Bronze Medal"
		default:
			answer[idx] = strconv.Itoa(len(score) - i)
		}
	}

	return answer
}

// func findRelativeRanks(score []int) []string {
// 	indices := make([]int, len(score))
// 	for i := range indices {
// 		indices[i] = i
// 	}
//
// 	slices.SortFunc(indices, func(i, j int) int {
// 		return score[j] - score[i]
// 	})
//
// 	answer := make([]string, len(score))
//
// 	for rankIdx, scoreIdx := range indices {
// 		rank := rankIdx + 1
//
// 		switch rank {
// 		case 1:
// 			answer[scoreIdx] = "Gold Medal"
// 		case 2:
// 			answer[scoreIdx] = "Silver Medal"
// 		case 3:
// 			answer[scoreIdx] = "Bronze Medal"
// 		default:
// 			answer[scoreIdx] = strconv.Itoa(rank)
// 		}
// 	}
//
// 	return answer
// }
