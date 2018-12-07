package p00028_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/weineel/go-sprint/leetcode/p00028"
)

var _ = Describe("P00028", func() {
	It("input: 'hello', '', expect: 0", func() {
		Expect(StrStr("hello", "")).To(Equal(0))
	})

	It("input: 'hello', 'hello', expect: 0", func() {
		Expect(StrStr("hello", "hello")).To(Equal(0))
	})

	It("input: 'hello', 'll', expect: 2", func() {
		Expect(StrStr("hello", "ll")).To(Equal(2))
	})

	It("input: 'weineel', 'nee', expect: 3", func() {
		Expect(StrStr("weineel", "nee")).To(Equal(3))
	})
})
