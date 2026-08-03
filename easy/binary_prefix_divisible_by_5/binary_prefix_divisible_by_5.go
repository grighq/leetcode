package binaryprefixdivisibleby5

func prefixesDivBy5(nums []int) []bool {
	res := make([]bool, len(nums))

	sum := 0
	for i, bit := range nums {
		sum = (sum*2 + bit) % 5
		res[i] = sum == 0
	}

	return res
}
