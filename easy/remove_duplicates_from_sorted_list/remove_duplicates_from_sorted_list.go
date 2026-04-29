package removeduplicatesfromsortedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	current := head
	for current != nil && current.Next != nil {
		if current.Val == current.Next.Val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}
	return head
}

func deleteDuplicatesRecursion(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	if head.Val == head.Next.Val {
		head.Next = head.Next.Next
		return deleteDuplicatesRecursion(head)
	}
	head.Next = deleteDuplicatesRecursion(head.Next)
	return head
}
