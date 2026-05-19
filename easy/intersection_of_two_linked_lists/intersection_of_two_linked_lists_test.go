package intersectionoftwolinkedlists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var cases = []struct {
	aVals      []int
	bVals      []int
	commonVals []int
	wantVal    int
	wantNil    bool
}{
	{[]int{}, []int{}, nil, 0, true},
	{[]int{1}, []int{}, nil, 0, true},
	{[]int{}, []int{2}, nil, 0, true},
	{[]int{1, 2, 3}, []int{4, 5, 6}, nil, 0, true},
	{[]int{1, 2, 3}, []int{1, 2, 3}, nil, 0, true},

	{[]int{4, 1}, []int{5, 6, 1}, []int{8, 4, 5}, 8, false},
	{[]int{}, []int{1, 2}, []int{3}, 3, false},
	{[]int{1, 2}, []int{}, []int{3}, 3, false},
	{[]int{1}, []int{2, 3}, []int{4, 5}, 4, false},
}

func TestGetIntersectionNode(t *testing.T) {
	for _, c := range cases {
		headA, headB, intersection := makeIntersectingLists(c.aVals, c.bVals, c.commonVals)

		got := getIntersectionNode(headA, headB)

		if c.wantNil {
			assert.Nil(t, got, "case aVals=%v bVals=%v commonVals=%v", c.aVals, c.bVals, c.commonVals)
		} else {
			if assert.NotNil(t, got, "case aVals=%v bVals=%v commonVals=%v", c.aVals, c.bVals, c.commonVals) {
				assert.Equal(t, intersection, got, "case aVals=%v bVals=%v commonVals=%v", c.aVals, c.bVals, c.commonVals)
				assert.Equal(t, c.wantVal, got.Val)
			}
		}
	}
}

func makeIntersectingLists(aVals, bVals, commonVals []int) (headA, headB, intersection *ListNode) {
	var commonHead *ListNode
	if len(commonVals) > 0 {
		commonHead = &ListNode{Val: commonVals[0]}
		curr := commonHead
		for _, v := range commonVals[1:] {
			curr.Next = &ListNode{Val: v}
			curr = curr.Next
		}
	}

	headA = buildList(aVals)
	if headA == nil {
		headA = commonHead
		intersection = commonHead
	} else {
		tailA := tailOf(headA)
		tailA.Next = commonHead
		intersection = commonHead
	}

	headB = buildList(bVals)
	if headB == nil {
		headB = commonHead
		intersection = commonHead
	} else {
		tailB := tailOf(headB)
		tailB.Next = commonHead
		intersection = commonHead
	}

	return headA, headB, intersection
}

func buildList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	head := &ListNode{Val: vals[0]}
	curr := head
	for _, v := range vals[1:] {
		curr.Next = &ListNode{Val: v}
		curr = curr.Next
	}
	return head
}

func tailOf(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	for head.Next != nil {
		head = head.Next
	}
	return head
}
