package parser

import (
	"github.com/romainmenke/css/stylesheet"
	"github.com/romainmenke/css/tokenizer"
)

func consumeAtRule(s tokenStream, inputToken tokenizer.Token) interface{} {
	rule := stylesheet.AtRule{
		Name: inputToken.String(),
	}

	for {
		t := s.Next()
		if t == nil {
			return rule
		}

		switch token := t.(type) {
		case tokenizer.TokenSemicolon:
			return rule
		case tokenizer.TokenEOF:
			return rule
		case tokenizer.TokenCurlyBracketLeft:
			rule.Block = consumeSimpleBlock(s, token)
			return rule
		default:
			rule.Prelude = append(rule.Prelude, consumeComponentValue(s, token))
		}
	}
}

func consumeQualifiedRule(s tokenStream, inputToken tokenizer.Token) stylesheet.Rule {
	rule := stylesheet.QualifiedRule{}

	for {
		t := s.Next()
		if t == nil {
			return rule
		}

		switch token := t.(type) {
		case tokenizer.TokenSemicolon:
			return rule
		case tokenizer.TokenEOF:
			return rule
		case tokenizer.TokenCurlyBracketLeft:
			rule.Block = consumeSimpleBlock(s, token)
			return rule
		default:
			rule.Prelude = append(rule.Prelude, consumeComponentValue(s, token))
		}
	}
}
