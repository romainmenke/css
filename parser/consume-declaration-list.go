package parser

import (
	"github.com/romainmenke/css/stylesheet"
	"github.com/romainmenke/css/tokenizer"
)

func consumeDeclarationList(s tokenStream, inputToken tokenizer.Token) interface{} {
	list := stylesheet.DeclarationList{}

	for {
		t1 := s.Next()
		if t1 == nil {
			return list
		}

		switch t1.(type) {
		case tokenizer.TokenWhitespace, tokenizer.TokenSemicolon:
			continue
		case tokenizer.TokenEOF:
			return list
		case tokenizer.TokenAtKeyword:
			rule := consumeAtRule(s, t1)
			if rule != nil {
				list = append(list, rule)
			}

		case tokenizer.TokenIdent:
			decl := consumeDeclaration(pipeWithCondition(s, streamCondition(func(t tokenizer.Token) bool {
				switch t.(type) {
				case tokenizer.TokenEOF, tokenizer.TokenSemicolon:
					return false
				default:
					return true
				}
			})))
			if decl != nil {
				decl.SetImportant()
				decl.RemoveTrailingWhitespace()

				list = append(list, *decl)
			}

		default:
			consumeComponentValue(s, t1)
		}
	}
}
