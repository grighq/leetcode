package palindromelinkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	mid := getMid(head)
	rev := reverseList(mid)

	return compare(rev, head)
}

func getMid(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow.Next
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		tmp := curr.Next
		curr.Next = prev
		prev = curr
		curr = tmp
	}

	return prev
}

func compare(list1, list2 *ListNode) bool {
	for list1 != nil {
		if list1.Val != list2.Val {
			return false
		}

		list1 = list1.Next
		list2 = list2.Next
	}

	return true
}
