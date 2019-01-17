package p00074_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/weineel/go-sprint/leetcode/p00074"
)

var _ = Describe("P00074", func() {
	matrix := [][]int{
		[]int{1, 3, 5, 7},
		[]int{10, 11, 16, 20},
		[]int{23, 30, 34, 50},
	}
	It(`
	Input:
	matrix = [
		[1,   3,  5,  7],
		[10, 11, 16, 20],
		[23, 30, 34, 50]
	]
	target = 3
	Output: true
	`, func() {
		Expect(SearchA2DMatrix(matrix, 3)).Should(BeTrue())
	})

	It(`
	Input:
	matrix = [
		[1,   3,  5,  7],
		[10, 11, 16, 20],
		[23, 30, 34, 50]
	]
	target = 13
	Output: false
	`, func() {
		Expect(SearchA2DMatrix(matrix, 13)).Should(BeFalse())
	})
})
