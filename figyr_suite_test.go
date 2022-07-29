package figyr_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestFigyr(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Figyr Suite")
}

var _ = Describe("Figyr", func() {
	It("works as intended", func() {

	})
})
