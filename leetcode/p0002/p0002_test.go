package p0002_test

import (
	"fmt"
	"strconv"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/weineel/go-sprint/kit"
	. "github.com/weineel/go-sprint/leetcode/p0002"
)

func MakeList(vals []int) *ListNode {
	vals = kit.ReverseInt(vals)
	var pre *ListNode = nil
	for _, val := range vals {
		pre = &ListNode{Val: val, Next: pre}
	}
	return pre
}

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

var _ = Describe("P0002", func() {
	It("Add Two Numbers", func() {
		// PrintList(MakeList([]int{8, 3, 4, 5}))
		l1 := MakeList([]int{2, 4, 3})
		l2 := MakeList([]int{5, 6, 4})
		sum := AddTwoNumbers(l1, l2)
		PrintList(sum)

		Expect(
			CompareList(
				sum,
				MakeList([]int{7, 0, 8}),
			)).Should(BeTrue())
	})

	It("Add Two Numbers 不等长链表", func() {
		l1 := MakeList([]int{8, 3, 4, 5, 5})
		l2 := MakeList([]int{8, 3, 4, 5})
		sum := AddTwoNumbers(l1, l2)
		PrintList(sum)

		Expect(
			CompareList(
				sum,
				MakeList([]int{6, 7, 8, 0, 6}),
			)).Should(BeTrue())
	})
})
