package rangesumqueryimmutable

type NumArray struct {
	PrefSums []int
}

func Constructor(nums []int) NumArray {
	na := NumArray{}
	na.PrefSums = make([]int, len(nums)+1)
	for i := 1; i <= len(nums); i++ {
		na.PrefSums[i] = na.PrefSums[i-1] + nums[i-1]
	}

	return na
}

func (na *NumArray) SumRange(left, right int) int {
	return na.PrefSums[right+1] - na.PrefSums[left]
}
