package p00376_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/weineel/go-sprint/leetcode/p00376"
)

var _ = Describe("P00376", func() {
	It("Input: [1,7,4,9,2,5], Output: 6", func() {
		Expect(WiggleSubsequence([]int{1, 7, 4, 9, 2, 5})).Should(Equal(6))
	})
	It("Input: [1,7], Output: 6", func() {
		Expect(WiggleSubsequence([]int{1, 7})).Should(Equal(2))
	})
	It("Input: [7,7,7,7], Output: 1", func() {
		Expect(WiggleSubsequence([]int{7, 7, 7, 7})).Should(Equal(1))
	})
	It("Input: [0,0], Output: 1", func() {
		Expect(WiggleSubsequence([]int{0, 0})).Should(Equal(1))
	})
})
