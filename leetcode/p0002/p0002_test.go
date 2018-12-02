package p0002_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/weineel/go-sprint/kit"
	. "github.com/weineel/go-sprint/leetcode/p0002"
)

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
