package lookup_test

import (
	"fmt"
	"os"

	"github.com/mplewis/figyr/lookup"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var supportedExtensions = map[string]string{
	".json": json,
	".yaml": yaml,
}

const json = `
{"dev_mode": true, "queue_size": 1000}
`

// supports any key casing style
const yaml = `
DevMode: true
QueueSize: 1000
`

var _ = Describe("NewFromFile", func() {
	testExt := func(ext string, raw string) {
		f, err := os.CreateTemp("", fmt.Sprintf("*.%s", ext))
		Expect(err).NotTo(HaveOccurred())
		defer f.Close()
		f.WriteString(raw)

		v, err := lookup.NewFromFile(f.Name())
		Expect(err).NotTo(HaveOccurred())

		val, found := v.Get("dev_mode")
		Expect(found).To(BeTrue())
		Expect(val).To(Equal("true"))

		val, found = v.Get("queue_size")
		Expect(found).To(BeTrue())
		Expect(val).To(Equal("1000"))
	}

	for ext, raw := range supportedExtensions {
		ext := ext
		raw := raw
		It(fmt.Sprintf("parses %s files", ext), func() {
			testExt(ext, raw)
		})
	}

	It("refuses to parse unsupported extensions", func() {
		_, err := lookup.NewFromFile("/tmp/bad-extension.txt")
		Expect(err).To(MatchError("unknown extension .txt"))
	})
})
