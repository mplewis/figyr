package lookup_test

import (
	"github.com/mplewis/figyr/lookup"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("args", func() {
	It("works as intended", func() {
		in := []string{
			"app.go",
			"--foo=bar", // only known arg format is --a=b
			"-test",
			"--baz=quux corge xyzzy", // supports quoted shell values
			"--xxx",
			"yyy",
			"some-pos-arg",
			"--foo=grault", // later overrides earlier
		}
		out := lookup.ValMap{"foo": "grault", "baz": "quux corge xyzzy"}
		Expect(lookup.NewFromArgs(in)).To(Equal(out))
	})
})
