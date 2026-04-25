package mergedtwosortedlists

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1, list2 *ListNode) *ListNode {
	dummy := new(ListNode)
	curr := dummy
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			curr.Next = list1
			list1 = list1.Next
		} else {
			curr.Next = list2
			list2 = list2.Next
		}
		curr = curr.Next
	}

	if list1 != nil {
		curr.Next = list1
	} else {
		curr.Next = list2
	}

	return dummy.Next
}

func mergeTwoListsRecursion(list1, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val <= list2.Val {
		list1.Next = mergeTwoListsRecursion(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoListsRecursion(list1, list2.Next)
		return list2
	}
}
