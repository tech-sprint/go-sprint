package p00086_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/weineel/go-sprint/kit"
	. "github.com/weineel/go-sprint/leetcode/p00086"
)

var _ = Describe("P00086", func() {
	It("{1, 4, 3, 2, 5, 2}, x = 3; expect: {1, 2, 2, 4, 3, 5}", func() {
		l := MakeList([]int{1, 4, 3, 2, 5, 2})
		expect := MakeList([]int{1, 2, 2, 4, 3, 5})

		result := PartitionList(l, 3)

		PrintList(result)

		Expect(
			CompareList(
				result,
				expect,
			)).Should(BeTrue())
	})

	It("{8, 3, 5, 2, 9}, x = 5; expect: {3, 2, 8, 5, 9}", func() {
		l := MakeList([]int{8, 3, 5, 2, 9})
		expect := MakeList([]int{3, 2, 8, 5, 9})

		result := PartitionList(l, 5)

		Expect(
			CompareList(
				result,
				expect,
			)).Should(BeTrue())
	})

	It("{8, 3, 5, 2, 9}, x = 8; expect: {3, 5, 2, 8, 9}", func() {
		l := MakeList([]int{8, 3, 5, 2, 9})
		expect := MakeList([]int{3, 5, 2, 8, 9})

		result := PartitionList(l, 8)

		Expect(
			CompareList(
				result,
				expect,
			)).Should(BeTrue())
	})

	It("{8, 3, 5, 2, 9}, x = 9; expect: {8, 3, 5, 2, 9}", func() {
		l := MakeList([]int{8, 3, 5, 2, 9})
		expect := MakeList([]int{8, 3, 5, 2, 9})

		result := PartitionList(l, 9)

		Expect(
			CompareList(
				result,
				expect,
			)).Should(BeTrue())
	})
})
