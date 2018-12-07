package p00028_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/weineel/go-sprint/leetcode/p00028"
)

var _ = Describe("P00028", func() {
	It("input: '', expect: 0", func() {
		Expect(StrStr("hello", "")).To(Equal(0))
	})
})
