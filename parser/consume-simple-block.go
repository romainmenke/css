package parser

import (
	"github.com/romainmenke/css/stylesheet"
	"github.com/romainmenke/css/tokenizer"
)

func consumeSimpleBlock(p *Parser, inputToken tokenizer.Token) stylesheet.SimpleBlock {
	simpleBlock := stylesheet.SimpleBlock{
		AssociatedToken: inputToken,
	}

	for {
		t := p.tz.Next()
		if t == nil {
			return simpleBlock
		}

		switch token := t.(type) {
		case tokenizer.TokenCurlyBracketRight:
			if token.IsMirror(inputToken) {
				return simpleBlock
			}
		case tokenizer.TokenSquareBracketRight:
			if token.IsMirror(inputToken) {
				return simpleBlock
			}
		case tokenizer.TokenParenthesisRight:
			if token.IsMirror(inputToken) {
				return simpleBlock
			}
		case tokenizer.TokenEOF:
			return simpleBlock
		default:
			simpleBlock.Value = append(simpleBlock.Value, consumeComponentValue(p, token))
		}
	}
}
