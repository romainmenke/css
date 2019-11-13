package parser

import (
	"github.com/romainmenke/css/stylesheet"
	"github.com/romainmenke/css/tokenizer"
)

func consumeAtRule(p *Parser, inputToken tokenizer.Token) stylesheet.Rule {
	rule := stylesheet.AtRule{
		Name: inputToken.String(),
	}

	for {
		t := p.tz.Next()
		if t == nil {
			return rule
		}

		switch token := t.(type) {
		case tokenizer.TokenSemicolon:
			return rule
		case tokenizer.TokenEOF:
			return rule
		case tokenizer.TokenCurlyBracketLeft:
			rule.Block = consumeSimpleBlock(p, token)
			return rule
		default:
			rule.Prelude = append(rule.Prelude, consumeComponentValue(p, token))
		}
	}
}

func consumeQualifiedRule(p *Parser, inputToken tokenizer.Token) stylesheet.Rule {
	rule := stylesheet.QualifiedRule{}

	for {
		t := p.tz.Next()
		if t == nil {
			return rule
		}

		switch token := t.(type) {
		case tokenizer.TokenSemicolon:
			return rule
		case tokenizer.TokenEOF:
			return nil
		case tokenizer.TokenCurlyBracketLeft:
			rule.Block = consumeSimpleBlock(p, token)
			return rule
		default:
			rule.Prelude = append(rule.Prelude, consumeComponentValue(p, token))
		}
	}
}
