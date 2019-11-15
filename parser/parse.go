package parser

import (
	"github.com/romainmenke/css/stylesheet"
)

func (p *Parser) ParseStylesheet() (stylesheet.Stylesheet, error) {
	return stylesheet.Stylesheet{
		Rules: consumeRuleList(p.tz, true),
	}, nil
}
