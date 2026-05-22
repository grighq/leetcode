package removelinkedlistelements

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {

	dummy := &ListNode{Next: head}
	prev, curr := dummy, head

	for curr != nil {
		if curr.Val == val {
			prev.Next = curr.Next
		} else {
			prev = prev.Next
		}
		curr = curr.Next
	}

	return dummy.Next
}

func removeElementsRecursion(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	head.Next = removeElementsRecursion(head.Next, val)
	if head.Val == val {
		return head.Next
	}

	return head
}
