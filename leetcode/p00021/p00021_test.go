package p00021_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/weineel/go-sprint/kit"
	. "github.com/weineel/go-sprint/leetcode/p00021"
)

var _ = Describe("P00021", func() {
	It("l1: {1, 2, 4}, l1: {1, 3, 4}; expect: {1, 1, 2, 3, 4, 4}", func() {
		l1 := MakeList([]int{1, 2, 4})
		l2 := MakeList([]int{1, 3, 4})
		expect := MakeList([]int{1, 1, 2, 3, 4, 4})

		result := MergeTwoSortedLists(l1, l2)

		PrintList(result)

		Expect(
			CompareList(
				result,
				expect,
			)).Should(BeTrue())
	})

	It("l1: {4, 7, 8}, l1: {8, 9, 10}; expect: {4, 7, 8, 8, 9, 10}", func() {
		l1 := MakeList([]int{4, 7, 8})
		l2 := MakeList([]int{8, 9, 10})
		expect := MakeList([]int{4, 7, 8, 8, 9, 10})

		result := MergeTwoSortedLists(l1, l2)

		PrintList(result)

		Expect(
			CompareList(
				result,
				expect,
			)).Should(BeTrue())
	})
})
