package figyr

import (
	"github.com/mplewis/figyr/lookup"
	"github.com/mplewis/figyr/refparse"
	"github.com/mplewis/figyr/types"
)

const defaultLineLength = 80

type Parser struct {
	types.ParserOptions
}

func New(description string) *Parser {
	return &Parser{types.ParserOptions{Description: description, LineLength: defaultLineLength}}
}

func (p *Parser) Parse(config any) error {
	source, err := lookup.NewFromDefaults(nil)
	if err != nil {
		return err
	}
	return refparse.Parse(p.ParserOptions, config, source)
}

func (p *Parser) MustParse(config any) {
	if err := p.Parse(config); err != nil {
		panic(err)
	}
}
