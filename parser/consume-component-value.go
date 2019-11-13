package parser

import (
	"github.com/romainmenke/css/stylesheet"
	"github.com/romainmenke/css/tokenizer"
)

func consumeComponentValue(p *Parser, inputToken tokenizer.Token) stylesheet.ComponentValue {
	for {
		t := p.tz.Next()
		if t == nil {
			return nil
		}

		switch token := t.(type) {
		case tokenizer.TokenCurlyBracketLeft, tokenizer.TokenSquareBracketLeft, tokenizer.TokenParenthesisLeft:
			return consumeSimpleBlock(p, token)
		case tokenizer.TokenFunction:
			return consumeFunction(p, token)
		default:
			return token
		}
	}
}
