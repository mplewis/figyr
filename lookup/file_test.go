package lookup_test

import (
	"fmt"
	"os"

	"github.com/mplewis/figyr/lookup"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

const json = `
{"foo": "bar", "baz": "quux"}
`

const yaml = `
foo: bar
baz: quux
`

const toml = `
foo = "bar"
baz = "quux"
`

var _ = Describe("NewFromFile", func() {
	It("works as intended", func() {
		testExt := func(ext string, raw string) {
			f, err := os.CreateTemp("", fmt.Sprintf("*.%s", ext))
			Expect(err).NotTo(HaveOccurred())
			defer f.Close()
			f.WriteString(raw)

			v, err := lookup.NewFromFile(f.Name())
			Expect(err).NotTo(HaveOccurred())
			Expect(v).To(Equal(lookup.ValMap{"foo": "bar", "baz": "quux"}))
		}

		raws := map[string]string{"json": json, "yaml": yaml, "toml": toml}
		for ext, raw := range raws {
			testExt(ext, raw)
		}

		_, err := lookup.NewFromFile("/tmp/bad-extension.txt")
		Expect(err).To(MatchError("unknown extension .txt"))
	})
})
