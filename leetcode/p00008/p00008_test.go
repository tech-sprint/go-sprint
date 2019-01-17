package p00008_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/weineel/go-sprint/kit"
	. "github.com/weineel/go-sprint/leetcode/p00008"
)

var _ = Describe("P00008", func() {
	fmt.Println(INT32MAX)
	fmt.Println(INT32MIN)

	Describe("Normal Situation", func() {
		It(`Input: "42", Output: 42`, func() {
			Expect(StringtoInteger("42")).Should(Equal(42))
		})

		It(`Input: "   -42", Output: -42`, func() {
			Expect(StringtoInteger("   -42")).Should(Equal(-42))
		})

		It(`Input: "    42", Output: 42`, func() {
			Expect(StringtoInteger("    42")).Should(Equal(42))
		})

		It(`Input: "4193 with words", Output: 4193`, func() {
			Expect(StringtoInteger("4193 with words")).Should(Equal(4193))
		})

		It(`Input: "  0000000000012345678", Output: 12345678`, func() {
			Expect(StringtoInteger("  0000000000012345678")).Should(Equal(12345678))
		})

		It(`Input: "  0000000000012345678sdf", Output: 12345678`, func() {
			Expect(StringtoInteger("  0000000000012345678sdf")).Should(Equal(12345678))
		})

		It(`Input: "9223372036854775808", Output: INT32MAX`, func() {
			Expect(StringtoInteger("9223372036854775808")).Should(Equal(INT32MAX))
		})

		It(`Input: "-91283472332", Output: INT32MIN`, func() {
			Expect(StringtoInteger("-91283472332")).Should(Equal(INT32MIN))
		})
		It(`Input: "-2147483649", Output: INT32MIN`, func() {
			Expect(StringtoInteger("-2147483649")).Should(Equal(INT32MIN))
		})

		It(`Input: "91283472332", Output: INT32MAX`, func() {
			Expect(StringtoInteger("91283472332")).Should(Equal(INT32MAX))
		})
		It(`Input: "2147483648", Output: INT32MAX`, func() {
			Expect(StringtoInteger("2147483648")).Should(Equal(INT32MAX))
		})
		It(`Input: "20000000000000000000", Output: INT32MAX`, func() {
			Expect(StringtoInteger("20000000000000000000")).Should(Equal(INT32MAX))
		})
	})

	Describe("Exception Situation", func() {
		It(`Input: "words and 987", Output: 0`, func() {
			Expect(StringtoInteger("words and 987")).Should(Equal(0))
		})

		It(`Input: "", Output: 0`, func() {
			Expect(StringtoInteger("")).Should(Equal(0))
		})

		It(`Input: " ", Output: 0`, func() {
			Expect(StringtoInteger(" ")).Should(Equal(0))
		})
	})
})
