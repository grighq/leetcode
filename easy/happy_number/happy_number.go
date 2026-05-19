package happynumber

func isHappy(n int) bool {

	getNext := func(num int) int {
		sum := 0
		for num != 0 {
			sum += (num % 10) * (num % 10)
			num /= 10
		}
		return sum
	}

	slow, fast := getNext(n), getNext(getNext(n))
	for slow != fast {
		slow, fast = getNext(slow), getNext(getNext(fast))
	}

	return slow == 1
}
