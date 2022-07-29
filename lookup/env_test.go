package lookup_test

import (
	"github.com/mplewis/figyr/lookup"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("NewFromEnv", func() {
	It("parses config values from the environment", func() {
		fakeEnv := map[string]string{
			"CHECK_INTERVAL":  "0.02",
			"SUCCESS_MESSAGE": "Your application was submitted",
			"someOtherEnvVar": "someValue", // only supports SCREAMING_SNAKE_CASE
		}
		fetcher := func(key string) string {
			return fakeEnv[key]
		}
		cfg := lookup.NewFromEnv(fetcher)

		val, found := cfg.Get("check_interval")
		Expect(found).To(BeTrue())
		Expect(val).To(Equal("0.02"))

		val, found = cfg.Get("success_message")
		Expect(found).To(BeTrue())
		Expect(val).To(Equal("Your application was submitted"))

		_, found = cfg.Get("some_other_env_var")
		Expect(found).To(BeFalse())
	})
})
