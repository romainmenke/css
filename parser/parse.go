package parser

import (
	"fmt"

	"github.com/romainmenke/css/stylesheet"
	"github.com/romainmenke/css/tokenizer"
)

func (p *Parser) Parse() (stylesheet.Stylesheet, error) {
	for {
		t := p.tz.Next()
		if _, ok := t.(tokenizer.TokenEOF); ok || t == nil {
			return stylesheet.Stylesheet{}, nil
		}

		switch token := t.(type) {
		case tokenizer.TokenWhitespace:
			continue
		case tokenizer.TokenAtKeyword:
			_ = consumeAtRule(p, token)
		default:
			fmt.Println(token)
		}

		break
	}

	return stylesheet.Stylesheet{}, nil
}
