package p00086

import (
	. "github.com/weineel/go-sprint/kit"
)

// PartitionList ...
func PartitionList(head *ListNode, x int) *ListNode {
	less := &ListNode{}
	lessTail := less
	more := &ListNode{}
	moreTail := more

	for head != nil {
		if head.Val < x {
			lessTail.Next = head
			lessTail = head
		} else {
			moreTail.Next = head
			moreTail = head
		}
		head = head.Next
	}
	lessTail.Next = more.Next
	moreTail.Next = nil
	return less.Next
}
