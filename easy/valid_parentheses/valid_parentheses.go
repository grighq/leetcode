package validparentheses

func isValid(s string) bool {
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	stack := make([]byte, 0, len(s))

	for i := range s {
		char := s[i]
		if open, isClose := pairs[char]; isClose {
			if len(stack) == 0 || stack[len(stack)-1] != open {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, char)
		}
	}

	return len(stack) == 0
}
