package twosum_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/weineel/go-sprint/leetcode/twosum"
)

var _ = Describe("TwoSum", func() {
	var commonInput []int

	BeforeEach(func() {
		commonInput = []int{1, 1, 4, 5, 7, 4}
	})

	Context("Have result", func() {
		It("Given an array of integers, return indices of the two numbers such that they add up to a specific target.", func() {
			Expect(TwoSum(commonInput, 5)).To(Or(Equal([]int{0, 2}), Equal([]int{1, 2})))
			Expect(TwoSum(commonInput, 11)).To(Or(Equal([]int{2, 4}), Equal([]int{4, 5})))
		})
		It("Should not use the same element twice.", func() {
			Expect(TwoSum(commonInput, 2)).NotTo(Equal([]int{0, 0}))
			Expect(TwoSum(commonInput, 2)).To(Equal([]int{0, 1}))

			Expect(TwoSum(commonInput, 8)).NotTo(Equal([]int{2, 2}))
			Expect(TwoSum(commonInput, 8)).To(Or(Equal([]int{0, 4}), Equal([]int{1, 4}), Equal([]int{2, 5})))
		})
	})

	It("No result, should return 'nil'", func() {
		Expect(TwoSum(commonInput, 1)).To(BeZero())
		Expect(TwoSum(commonInput, 100)).To(BeZero())
	})
})
