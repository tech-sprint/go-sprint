package p0002

import (
	. "github.com/weineel/go-sprint/kit"
)

// AddTwoNumbers 2. Add Two Numbers
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	cover := 0
	preSum := &ListNode{}
	sum := preSum
	for l1 != nil || l2 != nil || cover > 0 {
		a := 0
		b := 0
		if l1 != nil {
			a = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			b = l2.Val
			l2 = l2.Next
		}
		v := a + b + cover
		cover = v / 10
		sum.Next = &ListNode{Val: v % 10}
		sum = sum.Next
	}
	return preSum.Next
}
