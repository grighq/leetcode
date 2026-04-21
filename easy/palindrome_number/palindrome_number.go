package palindromenumber

func isPalindrome(num int) bool {
	if num < 0 {
		return false
	}
	tmpNum := num
	reverseNum := 0
	for tmpNum != 0 {
		reverseNum = reverseNum*10 + tmpNum%10
		tmpNum /= 10
	}
	return num == reverseNum
}
