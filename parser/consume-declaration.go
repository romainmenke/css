package parser

import (
	"github.com/romainmenke/css/stylesheet"
	"github.com/romainmenke/css/tokenizer"
)

func consumeDeclaration(p *Parser, inputToken tokenizer.Token) stylesheet.Declaration {
	declaration := stylesheet.Declaration{
		Name: inputToken.String(),
	}

	for {
		t1 := p.tz.Next()
		if t1 == nil {
			return nil
		}

		switch t1.(type) {
		case tokenizer.TokenWhitespace:
			continue
		case tokenizer.TokenColon:

			for {
				t2 := p.tz.Next()
				if t2 == nil {
					return nil
				}

				switch token2 := t2.(type) {
				case tokenizer.TokenEOF:
					return declaration

				default:
					declaration.Value = append(declaration.Value, consumeComponentValue(p, token2))
				}
			}

		default:
			return nil
		}
	}
}
