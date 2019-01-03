package p00012_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/weineel/go-sprint/leetcode/p00012"
)

var _ = Describe("P00012", func() {
	It("Input: 3, Output: \"III\"", func() {
		Expect(IntToRoman(3)).Should(Equal("III"))
	})

	It("Input: 4, Output: \"IV\"", func() {
		Expect(IntToRoman(4)).Should(Equal("IV"))
	})

	It("Input: 9, Output: \"IX\"", func() {
		Expect(IntToRoman(9)).Should(Equal("IX"))
	})

	It("Input: 58, Output: \"LVIII\"", func() {
		Expect(IntToRoman(58)).Should(Equal("LVIII"))
	})

	It("Input: 1994, Output: \"MCMXCIV\"", func() {
		Expect(IntToRoman(1994)).Should(Equal("MCMXCIV"))
	})

	It("Input: 3999, Output: \"MMMCMXCIX\"", func() {
		Expect(IntToRoman(3999)).Should(Equal("MMMCMXCIX"))
	})
})
