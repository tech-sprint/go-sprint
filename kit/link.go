package kit

import (
	"fmt"
	"strconv"
)

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// MakeList 生成一个链表
func MakeList(vals []int) *ListNode {
	vals = ReverseInt(vals)
	var pre *ListNode
	for _, val := range vals {
		pre = &ListNode{Val: val, Next: pre}
	}
	return pre
}

// PrintList 将链表打印到控制台
func PrintList(l *ListNode) {
	s := ""
	sep := ""
	for l != nil {
		s += sep + strconv.Itoa(l.Val)
		sep = "->"
		l = l.Next
	}
	fmt.Println(s)
}

// CompareList 比较两个链表的所有节点值是否相等。
func CompareList(l1 *ListNode, l2 *ListNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return l1 == nil && l2 == nil
}
