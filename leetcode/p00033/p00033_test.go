package p00033_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/weineel/go-sprint/leetcode/p00033"
)

var _ = Describe("P00033", func() {
	It("Input: nums = [], target = 7; Output: -1", func() {
		Expect(SearchInRotatedSortedArray([]int{}, 7)).Should(Equal(-1))
	})
	It("Input: nums = [7], target = 7; Output: 0", func() {
		Expect(SearchInRotatedSortedArray([]int{7}, 7)).Should(Equal(0))
	})
	It("Input: nums = [7], target = 5; Output: -1", func() {
		Expect(SearchInRotatedSortedArray([]int{7}, 5)).Should(Equal(-1))
	})
	It("Input: nums = [1,3], target = 3; Output: 1", func() {
		Expect(SearchInRotatedSortedArray([]int{1, 3}, 3)).Should(Equal(1))
	})
	It("Input: nums = [1,3], target = 3; Output: 0", func() {
		Expect(SearchInRotatedSortedArray([]int{1, 3}, 1)).Should(Equal(0))
	})

	It("Input: nums = [0,1,2,4,5,6,7], target = 5; Output: 4", func() {
		Expect(SearchInRotatedSortedArray([]int{0, 1, 2, 4, 5, 6, 7}, 5)).Should(Equal(4))
	})

	It("Input: nums = [4,5,6,7,0,1,2], target = 0; Output: 4", func() {
		Expect(SearchInRotatedSortedArray([]int{4, 5, 6, 7, 0, 1, 2}, 0)).Should(Equal(4))
	})

	It("Input: nums = [4,5,6,7,0,1,2], target = 1; Output: 5", func() {
		Expect(SearchInRotatedSortedArray([]int{4, 5, 6, 7, 0, 1, 2}, 1)).Should(Equal(5))
	})

	It("Input: nums = [4,5,6,7,0,1,2], target = 6; Output: 2", func() {
		Expect(SearchInRotatedSortedArray([]int{4, 5, 6, 7, 0, 1, 2}, 6)).Should(Equal(2))
	})

	It("Input: nums = [4,5,6,7,0,1,2], target = 9; Output: -1", func() {
		Expect(SearchInRotatedSortedArray([]int{4, 5, 6, 7, 0, 1, 2}, 9)).Should(Equal(-1))
	})
})
