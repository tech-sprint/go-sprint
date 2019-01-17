package p00704_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/weineel/go-sprint/leetcode/p00704"
)

var _ = Describe("P00704", func() {
	It("Input: nums = [], target = 7; Output: -1", func() {
		Expect(BinarySearch([]int{}, 7)).Should(Equal(-1))
	})
	It("Input: nums = [7], target = 7; Output: 0", func() {
		Expect(BinarySearch([]int{7}, 7)).Should(Equal(0))
	})
	It("Input: nums = [7], target = 5; Output: -1", func() {
		Expect(BinarySearch([]int{7}, 5)).Should(Equal(-1))
	})
	It("Input: nums = [1,3], target = 3; Output: 1", func() {
		Expect(BinarySearch([]int{1, 3}, 3)).Should(Equal(1))
	})
	It("Input: nums = [1,3], target = 3; Output: 0", func() {
		Expect(BinarySearch([]int{1, 3}, 1)).Should(Equal(0))
	})

	It("Input: nums = [0,1,2,4,5,6,7], target = 5; Output: 4", func() {
		Expect(BinarySearch([]int{0, 1, 2, 4, 5, 6, 7}, 5)).Should(Equal(4))
	})

	It("Input: nums = [-1,0,3,5,9,12], target = 9; Output: 4", func() {
		Expect(BinarySearch([]int{-1, 0, 3, 5, 9, 12}, 9)).Should(Equal(4))
	})

	It("Input: nums = [4,5,6,7,0,1,2], target = 2; Output: -1", func() {
		Expect(BinarySearch([]int{-1, 0, 3, 5, 9, 12}, 2)).Should(Equal(-1))
	})
})
