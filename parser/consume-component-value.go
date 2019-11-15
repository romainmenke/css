package parser

import (
	"github.com/romainmenke/css/tokenizer"
)

func consumeComponentValue(s tokenStream, reconsumed tokenizer.Token) interface{} {
	var t tokenizer.Token
	if reconsumed != nil {
		t = reconsumed
	} else {
		t = s.Next()
	}

	switch token := t.(type) {
	case tokenizer.TokenCurlyBracketLeft, tokenizer.TokenSquareBracketLeft, tokenizer.TokenParenthesisLeft:
		return consumeSimpleBlock(s, token)
	case tokenizer.TokenFunction:
		return consumeFunction(s, token)
	default:
		return token
	}
}
