package p00053_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/weineel/go-sprint/leetcode/p00053"
)

var _ = Describe("P00053", func() {

	a := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	a1 := 6

	b := []int{-2, 1, 2, 4, -1, 2, 1, -5, 4}
	b1 := 9

	c := []int{-2, 1}
	c1 := 1

	Describe("Only Maximum Value", func() {
		It("Input: {-2, 1, -3, 4, -1, 2, 1, -5, 4}, Output: 6.", func() {
			Expect(MaximumValue(a)).To(Equal(a1))
		})

		It("Input: {-2, 1, 2, 4, -1, 2, 1, -5, 4}, Output: 9.", func() {
			Expect(MaximumValue(b)).To(Equal(b1))
		})

		It("Input: {-2, 1}, Output: 1.", func() {
			Expect(MaximumValue(c)).To(Equal(c1))
		})
	})

	Describe("Maximum Value and Maximum Subarray", func() {
		It("Input: {-2, 1, -3, 4, -1, 2, 1, -5, 4}, Output: 6. [4, -1, 2, 1] has the largest sum = 6.", func() {
			Expect(MaximumSubarray(a)).To(Equal([3]int{6, 3, 6}))
		})

		It("Input: {-2, 1, 2, 4, -1, 2, 1, -5, 4}, Output: 9. [1, 2, 4, -1, 2, 1] has the largest sum = 9.", func() {
			Expect(MaximumSubarray(b)).To(Equal([3]int{9, 1, 6}))
		})

		It("Input: {-2, 1}, Output: 1. [1] has the largest sum = 1.", func() {
			Expect(MaximumSubarray(c)).To(Equal([3]int{1, 1, 1}))
		})
	})
})
