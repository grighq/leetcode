package countingbits

func countBits(n int) []int {
	bits := make([]int, n+1)
	for i := 1; i <= n; i++ {
		bits[i] = bits[i>>1] + i&1
		// bits[i] = bits[i&(i-1)] + 1
	}
	return bits
}
