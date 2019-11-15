package parser

import (
	"github.com/romainmenke/css/stylesheet"
	"github.com/romainmenke/css/tokenizer"
)

func consumeFunction(s tokenStream, reconsumed tokenizer.Token) interface{} {
	function := stylesheet.Function{
		Name: reconsumed.String(),
	}

	for {
		t := s.Next()
		if t == nil {
			return function
		}

		switch token := t.(type) {
		case tokenizer.TokenParenthesisRight:
			return function
		case tokenizer.TokenEOF:
			return function
		default:
			function.Value = append(function.Value, consumeComponentValue(s, token))
		}
	}
}
