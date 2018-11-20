package p0002_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestP0002(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "P0002 Suite")
}
