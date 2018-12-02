package p00021

import (
	. "github.com/weineel/go-sprint/kit"
)

// MergeTwoSortedLists ...
func MergeTwoSortedLists(l1 *ListNode, l2 *ListNode) *ListNode {
	temp := &ListNode{}
	tempTail := temp
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			tempTail.Next = l1
			l1 = l1.Next
		} else {
			tempTail.Next = l2
			l2 = l2.Next
		}
		tempTail = tempTail.Next
	}
	// 连接剩下的节点
	if l1 != nil {
		tempTail.Next = l1
	}
	if l2 != nil {
		tempTail.Next = l2
	}
	return temp.Next
}
