package parser

import (
	"github.com/romainmenke/css/stylesheet"
	"github.com/romainmenke/css/tokenizer"
)

func consumeSimpleBlock(s tokenStream, associatedToken tokenizer.Token) stylesheet.Block {
	simpleBlock := stylesheet.SimpleBlock{
		AssociatedToken: associatedToken,
	}

	for {
		t := s.Next()
		if t == nil {
			return simpleBlock
		}

		switch token := t.(type) {
		case tokenizer.MirroreableToken:
			if token.IsMirror(associatedToken) {
				return simpleBlock
			}

			simpleBlock.Value = append(simpleBlock.Value, consumeComponentValue(s, token))
		case tokenizer.TokenEOF:
			return simpleBlock
		default:
			simpleBlock.Value = append(simpleBlock.Value, consumeComponentValue(s, token))
		}
	}
}
