package refparse_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestRefparse(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Refparse Suite")
}
