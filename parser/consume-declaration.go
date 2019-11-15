package parser

import (
	"github.com/romainmenke/css/stylesheet"
	"github.com/romainmenke/css/tokenizer"
)

func consumeDeclaration(s tokenStream) *stylesheet.Declaration {
	nameToken := s.Next()
	if nameToken == nil {
		return nil
	}

	if _, ok := nameToken.(tokenizer.TokenEOF); ok {
		return nil
	}

	declaration := &stylesheet.Declaration{
		Name: nameToken.String(),
	}

	for {
		t1 := s.Next()
		if t1 == nil {
			return declaration
		}

		switch t1.(type) {
		case tokenizer.TokenWhitespace:
			continue
		case tokenizer.TokenColon:

			for {
				t2 := s.Next()
				if t2 == nil {
					return declaration
				}

				switch token2 := t2.(type) {
				case tokenizer.TokenEOF:
					return declaration

				default:
					declaration.Value = append(declaration.Value, consumeComponentValue(s, token2))
				}
			}

		default:
			return declaration
		}
	}
}
