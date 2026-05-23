package convertbinarynumberinalinkedlisttointeger

type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	res := 0
	for head != nil {
		// res = (res << 1) | head.Val
		res = res*2 + head.Val
		head = head.Next
	}

	return res
}
