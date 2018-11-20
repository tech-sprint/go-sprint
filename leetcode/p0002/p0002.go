package p0002

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// AddTwoNumbers 2. Add Two Numbers
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	next1 := l1
	next2 := l2
	cover := 0
	preSum := &ListNode{}
	sum := preSum
	for next1 != nil || next2 != nil || cover > 0 {
		a := 0
		b := 0
		if next1 != nil {
			a = next1.Val
			next1 = next1.Next
		}
		if next2 != nil {
			b = next2.Val
			next2 = next2.Next
		}
		v := a + b + cover
		cover = v / 10
		sum.Next = &ListNode{Val: v % 10}
		sum = sum.Next
	}
	return preSum.Next
}
