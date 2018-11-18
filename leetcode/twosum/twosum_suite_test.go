package twosum_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestTwosum(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Twosum Suite")
}
