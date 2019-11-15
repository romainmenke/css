package parser

import (
	"github.com/romainmenke/css/stylesheet"
	"github.com/romainmenke/css/tokenizer"
)

func consumeRuleList(s tokenStream, topLevel bool) []interface{} {
	list := stylesheet.RuleList{}

	for {
		t := s.Next()
		if t == nil {
			return list
		}

		switch t.(type) {
		case tokenizer.TokenWhitespace, tokenizer.TokenSemicolon:
			continue
		case tokenizer.TokenEOF:
			return list
		case tokenizer.TokenCDO, tokenizer.TokenCDC:
			if topLevel {
				continue
			}

			rule := consumeQualifiedRule(s, t)
			if rule != nil {
				list = append(list, rule)
			}
		case tokenizer.TokenAtKeyword:
			rule := consumeAtRule(s, t)
			if rule != nil {
				list = append(list, rule)
			}

		default:
			rule := consumeQualifiedRule(s, t)
			if rule != nil {
				list = append(list, rule)
			}
		}
	}
}
