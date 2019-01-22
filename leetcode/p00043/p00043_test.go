package p00043_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/weineel/go-sprint/leetcode/p00043"
)

var _ = Describe("P00043", func() {
	Describe("ADD", func() {
		It(`Input: num1 = "2", num2 = "3", Output: "5"`, func() {
			Expect(AddStrings("2", "3")).Should(Equal("5"))
		})

		It(`Input: num1 = "123", num2 = "456", Output: "579"`, func() {
			Expect(AddStrings("123", "456")).Should(Equal("579"))
		})

		It(`Input: num1 = "123", num2 = "4598", Output: "4721"`, func() {
			Expect(AddStrings("123", "4598")).Should(Equal("4721"))
		})

		It(`Input: num1 = "123", num2 = "19598", Output: "19721"`, func() {
			Expect(AddStrings("123", "19598")).Should(Equal("19721"))
		})

		It(`Input: num1 = "8923", num2 = "8923", Output: "17846"`, func() {
			Expect(AddStrings("8923", "8923")).Should(Equal("17846"))
		})
	})

	Describe("Multiply", func() {
		It(`Input: num1 = "2", num2 = "3", Output: "6"`, func() {
			Expect(MultiplyStrings("2", "3")).Should(Equal("6"))
		})

		It(`Input: num1 = "123", num2 = "456", Output: "56088"`, func() {
			Expect(MultiplyStrings("123", "456")).Should(Equal("56088"))
		})

		It(`Input: num1 = "123", num2 = "4598", Output: "565554"`, func() {
			Expect(MultiplyStrings("123", "4598")).Should(Equal("565554"))
		})

		It(`Input: num1 = "123", num2 = "19598", Output: "2410554"`, func() {
			Expect(MultiplyStrings("123", "19598")).Should(Equal("2410554"))
		})

		It(`Input: num1 = "9133", num2 = "0", Output: "0"`, func() {
			Expect(MultiplyStrings("9133", "0")).Should(Equal("0"))
		})

		It(`Input: num1 = "9033", num2 = "0", Output: "0"`, func() {
			Expect(MultiplyStrings("9133", "0")).Should(Equal("0"))
		})

		It(`Input: num1 = "0", num2 = "0", Output: "0"`, func() {
			Expect(MultiplyStrings("9133", "0")).Should(Equal("0"))
		})

		It(`Input: num1 = "0", num2 = "2320", Output: "0"`, func() {
			Expect(MultiplyStrings("9133", "0")).Should(Equal("0"))
		})
	})
})
