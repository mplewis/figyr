package figyr

import (
	"github.com/mplewis/figyr/lookup"
	"github.com/mplewis/figyr/refparse"
)

func Parse(config any) error {
	source, err := lookup.NewFromDefaults(nil)
	if err != nil {
		return err
	}
	return refparse.Parse(config, source)
}
