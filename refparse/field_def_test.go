package refparse_test

import (
	"reflect"

	"github.com/mplewis/figyr/refparse"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("FieldDef", func() {
	Describe("BreakValueIntoMandates", func() {
		It("breaks the raw value into a map of mandates", func() {
			value := `required`
			Expect(refparse.BreakValueIntoMandates(value)).To(Equal(map[string]string{
				"required": "",
			}))

			value = `description=do the thing,default=foo\,bar\,baz`
			Expect(refparse.BreakValueIntoMandates(value)).To(Equal(map[string]string{
				"description": "do the thing",
				"default":     "foo,bar,baz",
			}))
		})
	})

	Describe("BuildFieldDef", func() {
		It("builds a FieldDef from the given field name, type, and Figyr tag", func() {
			name := "MyField"
			typ := reflect.TypeOf("")

			tagVal := `required,description=do the thing`
			def, err := refparse.BuildFieldDef(name, typ, tagVal)
			Expect(err).ToNot(HaveOccurred())
			Expect(def).To(Equal(refparse.FieldDef{
				Name:        name,
				Type:        typ,
				Required:    true,
				Default:     "",
				Description: "do the thing",
			}))

			tagVal = `optional,description=do the thing`
			def, err = refparse.BuildFieldDef(name, typ, tagVal)
			Expect(err).ToNot(HaveOccurred())
			Expect(def).To(Equal(refparse.FieldDef{
				Name:        name,
				Type:        typ,
				Required:    false,
				Default:     "",
				Description: "do the thing",
			}))

			tagVal = `default=foo\,bar\,baz,description=do the thing`
			def, err = refparse.BuildFieldDef(name, typ, tagVal)
			Expect(err).ToNot(HaveOccurred())
			Expect(def).To(Equal(refparse.FieldDef{
				Name:        name,
				Type:        typ,
				Required:    false,
				Default:     "foo,bar,baz",
				Description: "do the thing",
			}))
		})
	})
})
