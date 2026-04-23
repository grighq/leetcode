package romantointeger

func romanToInt(s string) int {
	romanNums := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	sum, curr, prev := 0, 0, 0
	for i := range s {
		curr = romanNums[s[i]]
		if curr > prev {
			sum += curr - 2*prev
		} else {
			sum += curr
		}
		prev = curr
	}
	return sum
}
