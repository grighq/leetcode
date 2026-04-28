package plusone

func plusOne(digits []int) []int {
	for end := len(digits) - 1; end >= 0; end-- {
		if digits[end] != 9 {
			digits[end]++
			return digits
		}
		digits[end] = 0
	}
	return append([]int{1}, digits...)
}
