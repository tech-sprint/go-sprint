package p00455_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/weineel/go-sprint/leetcode/p00455"
)

var _ = Describe("P00455", func() {
	It("Input: [2, 3, 1], [1, 1]; Output: 1", func() {
		Expect(AssignCookies([]int{2, 3, 1}, []int{1, 1})).Should(Equal(1))
	})

	It("Input: [1, 2], [3, 1, 2]; Output: 1", func() {
		Expect(AssignCookies([]int{1, 2}, []int{3, 1, 2})).Should(Equal(2))
	})
})
