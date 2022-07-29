package lookup_test

import (
	"github.com/mplewis/figyr/lookup"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Combine", func() {
	It("combines getters with proper fallback", func() {
		primary := lookup.NewFromArgs([]string{
			"--check-interval=0.02",
			"--success-message=Your application was submitted",
		})
		fallback := lookup.NewFromArgs([]string{
			"--check-interval=0.10",
			"--port=9999",
		})
		combined := lookup.Combine(primary, fallback)

		val, found := combined.Get("check_interval")
		Expect(found).To(BeTrue())
		Expect(val).To(Equal("0.02"))

		val, found = combined.Get("success_message")
		Expect(found).To(BeTrue())
		Expect(val).To(Equal("Your application was submitted"))

		val, found = combined.Get("port")
		Expect(found).To(BeTrue())
		Expect(val).To(Equal("9999"))

		_, found = combined.Get("xxx")
		Expect(found).To(BeFalse())
	})
})
