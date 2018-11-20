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
	next := l
	s := ""
	sep := ""
	for next != nil {
		s += sep + strconv.Itoa(next.Val)
		sep = "->"
		next = next.Next
	}
	fmt.Println(s)
}

func CompareList(l1 *ListNode, l2 *ListNode) bool {
	next1 := l1
	next2 := l2
	for next1 != nil && next2 != nil {
		if next1.Val != next2.Val {
			return false
		}
		next1 = next1.Next
		next2 = next2.Next
	}
	return next1 == nil && next2 == nil
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
