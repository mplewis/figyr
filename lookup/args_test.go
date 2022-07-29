package lookup_test

import (
	"github.com/mplewis/figyr/lookup"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("NewFromArgs", func() {
	It("parses config values from os.Args", func() {
		args := []string{
			"app.go",
			"--check-interval=0.02", // only known arg format is --a=b
			"-test",
			"--success-message=Your application was submitted", // supports quoted shell values
			"--xxx",
			"yyy",
			"some-pos-arg",
			"--check-interval=0.05", // later overrides earlier
		}
		cfg := lookup.NewFromArgs(args)

		_, found := cfg.Get("test")
		Expect(found).To(BeFalse())

		val, found := cfg.Get("check_interval")
		Expect(found).To(BeTrue())
		Expect(val).To(Equal("0.05"))

		val, found = cfg.Get("success_message")
		Expect(found).To(BeTrue())
		Expect(val).To(Equal("Your application was submitted"))
	})
})
