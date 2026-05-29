package summaryranges

import "fmt"

func summaryRanges(nums []int) []string {
	p := 0
	strs := []string{}
	for i := 1; i <= len(nums); i++ {
		if i != len(nums) && nums[i]-nums[i-1] == 1 {
			continue
		}

		if i-p == 1 {
			strs = append(strs, fmt.Sprintf("%d", nums[p]))
		} else {
			strs = append(strs, fmt.Sprintf("%d->%d", nums[p], nums[i-1]))
		}

		p = i
	}

	return strs
}
