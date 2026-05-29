package summaryranges

import (
	"strconv"
)

func summaryRanges(nums []int) []string {
	p := 0
	ranges := []string{}
	for i := 1; i <= len(nums); i++ {
		if i != len(nums) && nums[i] == nums[i-1]+1 {
			continue
		}

		if i-p == 1 {
			ranges = append(ranges, strconv.Itoa(nums[p]))
		} else {
			ranges = append(ranges, strconv.Itoa(nums[p])+"->"+strconv.Itoa(nums[i-1]))
		}

		p = i
	}

	return ranges
}
